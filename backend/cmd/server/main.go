package main

import (
	"log/slog"

	"github.com/RhykerWells/RK/backend/internal/config"
	"github.com/RhykerWells/RK/backend/internal/database"
	"github.com/RhykerWells/RK/backend/internal/log"
)

func main() {
	log.Init(slog.LevelInfo)

	log.Logger().Info(
		"Starting Records Keeper",
	)

	cfg := config.Load()
	database.Connect(cfg)
}