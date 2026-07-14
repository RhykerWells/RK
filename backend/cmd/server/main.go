package main

import (
	"fmt"
	"log/slog"

	"github.com/RhykerWells/RK/backend/internal/config"
	"github.com/RhykerWells/RK/backend/internal/database"
	"github.com/RhykerWells/RK/backend/internal/logger"
	"github.com/RhykerWells/RK/backend/internal/server"
)

var Version = "0.0.0"

func main() {
	logger.Init(slog.LevelInfo)

	logger.Logger().Info(
		fmt.Sprintf("Starting Records Keeper V%s", Version),
	)

	config.Load()
	database.Connect()

	s := server.NewServer()
	s.Start()
}
