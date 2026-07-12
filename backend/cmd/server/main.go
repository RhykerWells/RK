package main

import (
	"log/slog"
	"os"

	"github.com/RhykerWells/RK/backend/internal/logger"
)

func main() {
	log := logger.New(os.Stdout, slog.LevelInfo)

	log.Info(
		"Starting Records Keeper",
	)
}