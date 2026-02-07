package api

import (
	"context"
	"errors"

	"github.com/database64128/proxy-sharing-go/api/admin"
	"github.com/database64128/proxy-sharing-go/api/ops"
	"github.com/database64128/proxy-sharing-go/ent"
	"github.com/database64128/proxy-sharing-go/jsoncfg"
	fiberzap "github.com/gofiber/contrib/v3/zap"
	"github.com/gofiber/fiber/v3"
	fiberlog "github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/etag"
	"github.com/gofiber/fiber/v3/middleware/pprof"
	"github.com/gofiber/fiber/v3/middleware/static"
	"go.uber.org/zap"
)

// Config stores the configuration for the RESTful API.
type Config struct {
	// DebugPprof enables pprof endpoints for debugging and profiling.
	DebugPprof bool `json:"debugPprof,omitzero"`

	// EnableTrustedProxyCheck enables trusted proxy checks.
	EnableTrustedProxyCheck bool `json:"enableTrustedProxyCheck,omitzero"`

	// TrustedProxies is the list of trusted proxies.
	// This only takes effect if EnableTrustedProxyCheck is true.
	TrustedProxies []string `json:"trustedProxies,omitzero"`

	// ProxyHeader is the header used to determine the client's IP address.
	// If empty, the remote peer's address is used.
	ProxyHeader string `json:"proxyHeader,omitzero"`

	// ListenNetwork is the network type.
	ListenNetwork string `json:"listenNetwork,omitzero"`

	// ListenAddress is the address to listen on.
	ListenAddress string `json:"listenAddress"`

	// ListenMode optionally sets the file mode of the unix domain socket.
	// It must be an octal number in a string (e.g., "0660").
	ListenMode jsoncfg.FileMode `json:"listenMode,omitzero"`

	// CertFile is the path to the certificate file.
	// If empty, TLS is disabled.
	CertFile string `json:"certFile,omitzero"`

	// KeyFile is the path to the key file.
	// This is required if CertFile is set.
	KeyFile string `json:"keyFile,omitzero"`

	// ClientCertFile is the path to the client certificate file.
	// If empty, client certificate authentication is disabled.
	ClientCertFile string `json:"clientCertFile,omitzero"`

	// StaticPath is the path where static files are served from.
	// If empty, static file serving is disabled.
	StaticPath string `json:"staticPath,omitzero"`

	// SecretPath adds a secret path prefix to all routes.
	// If empty, no secret path is added.
	SecretPath string `json:"secretPath,omitzero"`

	// FiberConfigPath overrides the [fiber.Config] settings we use.
	// If empty, no overrides are applied.
	FiberConfigPath string `json:"fiberConfigPath,omitzero"`

	// Admin is the configuration for the admin API.
	Admin admin.Config `json:"admin,omitzero"`

	// OpenProxySharing is the configuration for the Open Proxy Sharing API.
	OpenProxySharing ops.Config `json:"openProxySharing,omitzero"`
}

// Server returns a new API server from the config.
func (c *Config) Server(logger *zap.Logger, client *ent.Client) (*Server, error) {
	fiberlog.SetLogger(fiberzap.NewLogger(fiberzap.LoggerConfig{
		SetLogger: logger,
	}))

	fc := fiber.Config{
		ProxyHeader: c.ProxyHeader,
		TrustProxy:  c.EnableTrustedProxyCheck,
		TrustProxyConfig: fiber.TrustProxyConfig{
			Proxies: c.TrustedProxies,
		},
	}

	if c.FiberConfigPath != "" {
		if err := jsoncfg.Open(c.FiberConfigPath, &fc); err != nil {
			return nil, err
		}
	}

	app := fiber.New(fc)

	app.Use(etag.New())

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))

	var router fiber.Router = app
	if c.SecretPath != "" {
		if c.SecretPath[0] != '/' {
			c.SecretPath = "/" + c.SecretPath
		}
		router = app.Group(c.SecretPath)
	}

	if c.DebugPprof {
		app.Use(pprof.New(pprof.Config{
			Prefix: c.SecretPath,
		}))
	}

	api := router.Group("/api")

	// /api/admin/v1
	c.Admin.RegisterRoutes(api.Group("/admin/v1"), client, logger)

	// /api/ops/v1
	c.OpenProxySharing.RegisterRoutes(api.Group("/ops/v1"), client, logger)

	if c.StaticPath != "" {
		router.Use(static.New(c.StaticPath, static.Config{
			ByteRange: true,
		}))
	}

	listenNetwork := c.ListenNetwork
	if listenNetwork == "" {
		listenNetwork = "tcp"
	}

	return &Server{
		logger:         logger,
		app:            app,
		listenNetwork:  listenNetwork,
		listenAddress:  c.ListenAddress,
		listenMode:     c.ListenMode,
		certFile:       c.CertFile,
		keyFile:        c.KeyFile,
		clientCertFile: c.ClientCertFile,
	}, nil
}

// Server is the RESTful API server.
type Server struct {
	logger         *zap.Logger
	app            *fiber.App
	listenNetwork  string
	listenAddress  string
	listenMode     jsoncfg.FileMode
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
		flc := fiber.ListenConfig{
			ListenerNetwork:       s.listenNetwork,
			CertFile:              s.certFile,
			CertKeyFile:           s.keyFile,
			CertClientFile:        s.clientCertFile,
			UnixSocketFileMode:    s.listenMode.Value(),
			DisableStartupMessage: true,
		}
		if err := s.app.Listen(s.listenAddress, flc); err != nil {
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
