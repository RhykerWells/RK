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
	loadSchemas bool
	all         bool
)

func init() {
	flag.BoolVar(&flagVersion, "version", false, "print version and exit")
	flag.BoolVar(&loadSchemas, "load-schemas", false, "load database schemas and exit")
	flag.BoolVar(&all, "all", false, "run application")
}

func Init() {
	logger.Init(slog.LevelInfo)

	if !flag.Parsed() {
		flag.Parse()
	}

	if !flagVersion && !loadSchemas && !all {
		fmt.Println("No flags provided. Use -h for help.")
		os.Exit(1)
	}

	if flagVersion {
		fmt.Printf("Records Keeper V%s\n", internal.Version)
		os.Exit(0)
	}

	if loadSchemas {
		config.Load()
		database.Connect()
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
