package api

import (
	"context"
	"errors"

	v1 "github.com/database64128/proxy-sharing-go/api/v1"
	"github.com/database64128/proxy-sharing-go/ent"
	"github.com/database64128/proxy-sharing-go/jsonhelper"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"go.uber.org/zap"
)

// Config stores the configuration for the RESTful API.
type Config struct {
	// Debug
	DebugPprof bool `json:"debugPprof"`

	// Reverse proxy
	EnableTrustedProxyCheck bool     `json:"enableTrustedProxyCheck"`
	TrustedProxies          []string `json:"trustedProxies"`
	ProxyHeader             string   `json:"proxyHeader"`

	// Listen
	ListenAddress  string `json:"listenAddress"`
	CertFile       string `json:"certFile"`
	KeyFile        string `json:"keyFile"`
	ClientCertFile string `json:"clientCertFile"`

	// Static
	StaticPath string `json:"staticPath"`

	// Misc
	SecretPath      string `json:"secretPath"`
	FiberConfigPath string `json:"fiberConfigPath"`
}

// Server returns a new API server from the config.
func (c *Config) Server(logger *zap.Logger, client *ent.Client) (*Server, error) {
	fiberlog.SetLogger(fiberzap.NewLogger(fiberzap.LoggerConfig{
		SetLogger: logger,
	}))

	fc := fiber.Config{
		ProxyHeader:             c.ProxyHeader,
		DisableStartupMessage:   true,
		Network:                 "tcp",
		EnableTrustedProxyCheck: c.EnableTrustedProxyCheck,
		TrustedProxies:          c.TrustedProxies,
	}

	if c.FiberConfigPath != "" {
		if err := jsonhelper.OpenAndDecodeDisallowUnknownFields(c.FiberConfigPath, &fc); err != nil {
			return nil, err
		}
	}

	app := fiber.New(fc)

	app.Use(etag.New())

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
		Fields: []string{"latency", "status", "method", "url", "ip"},
	}))

	if c.DebugPprof {
		app.Use(pprof.New())
	}

	var router fiber.Router = app
	if c.SecretPath != "" {
		router = app.Group(c.SecretPath)
	}

	api := router.Group("/api")

	// /api/v1
	v1.Register(logger, client, api.Group("/v1"))

	if c.StaticPath != "" {
		router.Static("/", c.StaticPath, fiber.Static{
			ByteRange: true,
		})
	}

	return &Server{
		logger:         logger,
		app:            app,
		listenAddress:  c.ListenAddress,
		certFile:       c.CertFile,
		keyFile:        c.KeyFile,
		clientCertFile: c.ClientCertFile,
	}, nil
}

// Server is the RESTful API server.
type Server struct {
	logger         *zap.Logger
	app            *fiber.App
	listenAddress  string
	certFile       string
	keyFile        string
	clientCertFile string
	ctx            context.Context
}

// String implements [service.Service.String].
func (*Server) String() string {
	return "API server"
}

// Start starts the API server.
func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("Starting API server", zap.String("listenAddress", s.listenAddress))
	s.ctx = ctx
	go func() {
		var err error
		switch {
		case s.clientCertFile != "":
			err = s.app.ListenMutualTLS(s.listenAddress, s.certFile, s.keyFile, s.clientCertFile)
		case s.certFile != "":
			err = s.app.ListenTLS(s.listenAddress, s.certFile, s.keyFile)
		default:
			err = s.app.Listen(s.listenAddress)
		}
		if err != nil {
			s.logger.Fatal("Failed to start API server", zap.Error(err))
		}
	}()
	return nil
}

// Stop stops the API server.
func (s *Server) Stop() error {
	if err := s.app.ShutdownWithContext(s.ctx); err != nil {
		if errors.Is(err, context.Canceled) {
			return nil
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil
		}
		return err
	}
	return nil
}
