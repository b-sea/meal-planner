// Package api implements the meal prepping API.
package api

import (
	"context"
	"os"
	"time"

	"github.com/b-sea/go-server/server"
	"github.com/b-sea/meal-planner/internal/dash"
	"github.com/b-sea/meal-planner/internal/mock"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

const defaultLogLevel = zerolog.InfoLevel

// Service is the main API service.
type Service struct {
	dash dash.DASH
	http *server.Server
}

// New creates a new API service with the default server.
func New(version string, config Config) *Service {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack //nolint: reassign
	zerolog.TimeFieldFormat = time.RFC3339Nano

	level, err := zerolog.ParseLevel(config.Logger.Level)
	if err != nil {
		level = defaultLogLevel
	}

	log := zerolog.New(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		},
	).Level(level).With().Timestamp().Logger()

	zerolog.DefaultContextLogger = &log

	return NewService(
		server.New(
			context.Background(),
			&mock.Recorder{},
			server.WithVersion(version),
			server.WithPort(config.HTTP.Port),
			server.WithReadTimeout(time.Duration(config.HTTP.ReadTimeout)*time.Second),
			server.WithReadTimeout(time.Duration(config.HTTP.WriteTimeout)*time.Second),
		),
	)
}

// NewService creates a new API service with a custom server.
func NewService(svr *server.Server) *Service {
	service := &Service{
		dash: dash.New(),
		http: svr,
	}
	service.handle()

	return service
}

// Start the API service.
func (s *Service) Start() error {
	ctx := context.Background()

	if err := s.http.Start(ctx); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("start server")

		return err //nolint: wrapcheck
	}

	return nil
}

// Stop the API service.
func (s *Service) Stop() error {
	ctx := context.Background()

	if err := s.http.Stop(ctx); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("stop server")

		return err //nolint: wrapcheck
	}

	zerolog.Ctx(ctx).Info().Msg("stop server")

	return nil
}

func (s *Service) handle() {
	s.http.Router().PathPrefix("/api").Subrouter()
}
