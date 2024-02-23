package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/database64128/proxy-sharing-go/ent"
	"github.com/database64128/proxy-sharing-go/ent/migrate"
	"github.com/database64128/proxy-sharing-go/jsonhelper"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

// Config is the configuration for the database.
type Config struct {
	// Driver is the database driver.
	Driver string `json:"driver"`

	// DSN is the data source name.
	DSN string `json:"dsn"`

	// MaxOpenConns is the maximum number of open connections to the database.
	MaxOpenConns int `json:"maxOpenConns"`

	// MaxIdleConns is the maximum number of connections in the idle connection pool.
	MaxIdleConns int `json:"maxIdleConns"`

	// ConnMaxLifetime is the maximum amount of time a connection may be reused.
	ConnMaxLifetime jsonhelper.Duration `json:"connMaxLifetime"`

	// ConnMaxIdleTime is the maximum amount of time a connection may be idle.
	ConnMaxIdleTime jsonhelper.Duration `json:"connMaxIdleTime"`

	// Debug enables verbose logging.
	Debug bool `json:"debug"`

	// NoAutoMigrate disables auto-migration.
	NoAutoMigrate bool `json:"noAutoMigrate"`
}

// Open opens the database and runs auto-migration.
func (c *Config) Open(ctx context.Context, logger *zap.Logger) (*ent.Client, error) {
	var (
		db  *sql.DB
		err error
	)

	switch c.Driver {
	case dialect.MySQL:
		db, err = sql.Open(c.Driver, c.DSN)
	case dialect.SQLite:
		db, err = openSQLiteDB(c.DSN)
	case dialect.Postgres:
		db, err = sql.Open("pgx", c.DSN)
	default:
		return nil, fmt.Errorf("unsupported driver: %q", c.Driver)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime))
	db.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTime))

	drv := entsql.OpenDB(c.Driver, db)
	var opts []ent.Option
	if !c.Debug {
		opts = []ent.Option{ent.Driver(drv)}
	} else {
		sugar := logger.Sugar()
		opts = []ent.Option{ent.Driver(drv), ent.Debug(), ent.Log(sugar.Debugln)}
	}
	client := ent.NewClient(opts...)

	if !c.NoAutoMigrate {
		if err = client.Schema.Create(
			ctx,
			migrate.WithDropColumn(true),
			migrate.WithDropIndex(true),
		); err != nil {
			return nil, fmt.Errorf("failed to create schema resources: %w", err)
		}
	}

	return client, nil
}
