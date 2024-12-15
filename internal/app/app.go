package app

import (
	"context"
	"fmt"
	"log/slog"
	appConfig "smart_school_for_mirea/internal/config"
	"smart_school_for_mirea/internal/controller"
	appCore "smart_school_for_mirea/internal/core"
	"smart_school_for_mirea/internal/storage"
)

type App struct {
	logger  *slog.Logger
	storage *storage.PgStorage
	core    *appCore.Core
	api     *controller.Controller
}

func NewApp(logger *slog.Logger) (*App, error) {
	config, err := appConfig.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("get config: %w", err)
	}

	pgStorage := storage.NewPgStorage(config.DbURL)

	core := appCore.NewCore(pgStorage, config.AuthSecretKey, config.AuthTTL, logger)
	api := controller.NewController(core, config.ApiServerPort, config.AuthSecretKey, logger)

	return &App{
		logger:  logger,
		storage: pgStorage,
		core:    core,
		api:     api,
	}, nil
}

func (p *App) Start() error {
	p.logger.Info("Starting app")
	err := p.storage.Connect(context.Background())
	if err != nil {
		return fmt.Errorf("connection to database: %w", err)
	}

	return p.api.Start()
}
