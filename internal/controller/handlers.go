package controller

import (
	"log/slog"
	"smart_school_for_mirea/internal/core"
)

type handlers struct {
	core *core.Core

	logger *slog.Logger
}

func newHandlers(core *core.Core, logger *slog.Logger) *handlers {
	return &handlers{
		core: core,

		logger: logger,
	}
}
