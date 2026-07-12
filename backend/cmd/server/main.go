package main

import (
	"log/slog"

	"github.com/RhykerWells/RK/backend/internal/config"
	"github.com/RhykerWells/RK/backend/internal/logger"
)

func main() {
	logger.Init(slog.LevelInfo)

	logger.Info(
		"Starting Records Keeper",
	)

	config.Load()
}