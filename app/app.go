package app

import (
	"context"
	"dbKursach/internal/config"
	"dbKursach/internal/pkg/httperrors"
	"dbKursach/internal/pkg/logger"

	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type App struct {
	fiber  *fiber.App
	cfg    *config.Config
	pgDB   *sqlx.DB
	logger logger.Logger
}

func NewApp(
	cfg *config.Config,
	database *sqlx.DB,
	logger logger.Logger,
) *App {
	return &App{
		fiber:  fiber.New(fiber.Config{ErrorHandler: httperrors.Init(cfg, logger)}),
		cfg:    cfg,
		pgDB:   database,
		logger: logger,
	}
}

func (s *App) Run(ctx context.Context) error {
	if err := s.MapHandlers(ctx, s.fiber); err != nil {
		s.logger.Fatalf("Cannot map delivery: ", err)
	}

	go func() {
		s.logger.Infof("Start server on port: %s:%s", s.cfg.Server.Host, s.cfg.Server.Port)
		if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)); err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	err := s.fiber.Shutdown()
	if err != nil {
		s.logger.Error(err)
	} else {
		s.logger.Info("Fiber server exited properly")
	}

	return nil
}
