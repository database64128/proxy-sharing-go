package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	proxysharinggo "github.com/database64128/proxy-sharing-go"
	"github.com/database64128/proxy-sharing-go/jsonhelper"
	"github.com/database64128/proxy-sharing-go/logging"
	"github.com/database64128/proxy-sharing-go/service"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	version  bool
	testConf bool
	confPath string
	zapConf  string
	logLevel zapcore.Level
)

func init() {
	flag.BoolVar(&version, "version", false, "Print the version and exit")
	flag.BoolVar(&testConf, "testConf", false, "Test the configuration file and exit without starting the services")
	flag.StringVar(&confPath, "confPath", "config.json", "Path to the JSON configuration file")
	flag.StringVar(&zapConf, "zapConf", "console", "Preset name or path to the JSON configuration file for building the zap logger.\nAvailable presets: console, console-nocolor, console-notime, systemd, production, development")
	flag.TextVar(&logLevel, "logLevel", zapcore.InfoLevel, "Log level for the console and systemd presets.\nAvailable levels: debug, info, warn, error, dpanic, panic, fatal")
}

func main() {
	flag.Parse()

	if version {
		os.Stdout.WriteString("proxy-sharing-go " + proxysharinggo.Version + "\n")
		if info, ok := debug.ReadBuildInfo(); ok {
			os.Stdout.WriteString(info.String())
		}
		return
	}

	logger, err := logging.NewZapLogger(zapConf, logLevel)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to build logger:", err)
		os.Exit(1)
	}
	defer logger.Sync()

	logger.Info("proxy-sharing-go", zap.String("version", proxysharinggo.Version))

	var sc service.Config
	if err = jsonhelper.OpenAndDecodeDisallowUnknownFields(confPath, &sc); err != nil {
		logger.Fatal("Failed to load config",
			zap.String("confPath", confPath),
			zap.Error(err),
		)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ctx.Done()
		logger.Info("Received exit signal")
		stop()
	}()

	m, err := sc.Manager(ctx, logger)
	if err != nil {
		logger.Fatal("Failed to create service manager",
			zap.String("confPath", confPath),
			zap.Error(err),
		)
	}
	defer m.Close()

	if testConf {
		logger.Info("Config test OK", zap.String("confPath", confPath))
		return
	}

	if err = m.Start(ctx); err != nil {
		logger.Fatal("Failed to start services",
			zap.String("confPath", confPath),
			zap.Error(err),
		)
	}

	<-ctx.Done()
	m.Stop()
}
