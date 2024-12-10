package main

import (
	"log/slog"
	"os"
	"smart_school_for_mirea/internal/app"
	logTools "smart_school_for_mirea/pkg/logger"
)

func main() {
	logger := logTools.NewLogger(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	application, err := app.NewApp(logger)
	if err != nil {
		logger.Error("App don't created", slog.Any("error", err))
		return
	}
	err = application.Start()
	if err != nil {
		logger.Error("App don't started", slog.Any("error", err))
	}
}
