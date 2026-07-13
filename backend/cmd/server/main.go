package main

import (
	"fmt"
	"log/slog"

	"github.com/RhykerWells/RK/backend/internal/config"
	"github.com/RhykerWells/RK/backend/internal/database"
	"github.com/RhykerWells/RK/backend/internal/log"
)

var Version = "0.0.0"

func main() {
	log.Init(slog.LevelInfo)

	log.Logger().Info(
		fmt.Sprintf("Starting Records Keeper V%s", Version),
	)

	cfg := config.Load()
	database.Connect(cfg)
}
