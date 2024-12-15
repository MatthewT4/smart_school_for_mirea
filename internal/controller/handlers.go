package controller

import (
	"log/slog"
	"smart_school_for_mirea/internal/core"
)

type handlers struct {
	core *core.Core

	authSecretKey string

	logger *slog.Logger
}
