package run

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/RhykerWells/RK/backend/internal"
	"github.com/RhykerWells/RK/backend/internal/config"
	"github.com/RhykerWells/RK/backend/internal/database"
	"github.com/RhykerWells/RK/backend/internal/logger"
	"github.com/RhykerWells/RK/backend/internal/server"
)

var (
	flagVersion bool
)

func init() {
	flag.BoolVar(&flagVersion, "version", false, "print version and exit")
}

func Init() {
	logger.Init(slog.LevelInfo)

	if !flag.Parsed() {
		flag.Parse()
	}

	if flagVersion {
		fmt.Printf("Records Keeper V%s\n", internal.Version)
		os.Exit(0)
	}

	logger.Info(fmt.Sprintf("Starting Records Keeper V%s", internal.Version))

	config.Load()

	database.Connect()
}

func Run() {
	s := server.NewServer()

	s.Start()
}
