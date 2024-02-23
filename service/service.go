package service

import (
	"context"
	"fmt"

	"github.com/database64128/proxy-sharing-go/api"
	"github.com/database64128/proxy-sharing-go/database"
	"github.com/database64128/proxy-sharing-go/ent"
	"go.uber.org/zap"
)

// Config is the main configuration structure.
// It may be marshaled as or unmarshaled from JSON.
type Config struct {
	API      api.Config      `json:"api"`
	DataBase database.Config `json:"database"`
}

// Manager initializes the service manager.
func (sc *Config) Manager(ctx context.Context, logger *zap.Logger) (*Manager, error) {
	client, err := sc.DataBase.Open(ctx, logger)
	if err != nil {
		return nil, err
	}

	apiServer, err := sc.API.Server(logger, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create API server: %w", err)
	}

	return &Manager{
		services: []Service{apiServer},
		logger:   logger,
		client:   client,
	}, nil
}

// Service implements the business logic.
type Service interface {
	// String returns the relay service's name.
	String() string

	// Start starts the relay service.
	Start(ctx context.Context) error

	// Stop stops the relay service.
	Stop() error
}

// Manager manages the services.
type Manager struct {
	services []Service
	logger   *zap.Logger
	client   *ent.Client
}

// Start starts all configured services.
func (m *Manager) Start(ctx context.Context) error {
	for _, s := range m.services {
		if err := s.Start(ctx); err != nil {
			return fmt.Errorf("failed to start %s: %w", s.String(), err)
		}
	}
	return nil
}

// Stop stops all running services.
func (m *Manager) Stop() {
	for _, s := range m.services {
		if err := s.Stop(); err != nil {
			m.logger.Warn("Failed to stop service",
				zap.Stringer("service", s),
				zap.Error(err),
			)
		}
		m.logger.Info("Stopped service", zap.Stringer("service", s))
	}
}

// Close closes the manager.
func (m *Manager) Close() {
	if err := m.client.Close(); err != nil {
		m.logger.Warn("Failed to close database connection", zap.Error(err))
	}
	m.logger.Info("Closed database connection")
}
