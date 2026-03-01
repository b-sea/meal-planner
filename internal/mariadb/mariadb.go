// Package mariadb implements all interactions with MariaDB.
package mariadb

import (
	"context"
	"database/sql"
	"embed"
	"path/filepath"
	"time"

	// Register the mysql driver.
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"

	"github.com/b-sea/go-server/server"
)

const defaultTimeout = 20 * time.Second

//go:embed sql/*.sql
var sqlFS embed.FS

// Recorder defines functions required for MariaDB metrics.
type Recorder interface {
	ObserveMariadbTxDuration(status string, duration time.Duration)
}

var _ server.HealthChecker = (*Repository)(nil)

// Repository implements all MariaDB interactions.
type Repository struct {
	db       *sql.DB
	recorder Recorder
	timeout  time.Duration
}

// New creates a new MariaDB repository with a standard connection.
func New(host string, user string, pwd string, recorder Recorder) *Repository {
	return NewRepository(connector(host, user, pwd), recorder)
}

// NewRepository creates a new MariaDB repository with a custom connection.
func NewRepository(connector Connector, recorder Recorder) *Repository {
	return &Repository{
		db:       connector(),
		recorder: recorder,
		timeout:  defaultTimeout,
	}
}

// HealthCheck checks the healthiness of MariaDB.
func (r *Repository) HealthCheck(ctx context.Context) error {
	if err := r.db.PingContext(ctx); err != nil {
		return databaseError(err)
	}

	return nil
}

// Setup the database schema.
func (r *Repository) Setup() error {
	var err error

	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	log := zerolog.Ctx(ctx)

	defer func() {
		event := log.Info() //nolint: zerologlint

		if err != nil {
			event = log.Error().Err(err) //nolint: zerologlint
		}

		event.Dur("duration_ms", time.Since(start)).Msg("database setup")
	}()

	entries, err := sqlFS.ReadDir("sql")
	if err != nil {
		return fileReadError(err)
	}

	err = r.withTx(ctx, func(tx *sql.Tx) error {
		for _, entry := range entries {
			file := filepath.Join("sql", entry.Name())

			log.Debug().Str("file", file).Msg("load file")

			cmd, err := sqlFS.ReadFile(file)
			if err != nil {
				return fileReadError(err)
			}

			if _, err := tx.ExecContext(ctx, string(cmd)); err != nil {
				return err //nolint: wrapcheck
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Close the current database connection.
func (r *Repository) Close() error {
	if err := r.db.Close(); err != nil {
		return databaseError(err)
	}

	return nil
}

func (r *Repository) withTx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	var err error

	start := time.Now()

	defer func() {
		log := zerolog.Ctx(ctx)

		event := log.Debug() //nolint: zerologlint
		status := "success"

		if err != nil {
			event = log.Error().Err(err) //nolint: zerologlint
			status = "failed"
		}

		duration := time.Since(start)

		event.Dur("duration_ms", duration).Msg("database tx")
		r.recorder.ObserveMariadbTxDuration(status, duration)
	}()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return transactionError(err)
	}

	if err = fn(tx); err != nil {
		_ = tx.Rollback()

		return transactionError(err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()

		return transactionError(err)
	}

	return nil
}
