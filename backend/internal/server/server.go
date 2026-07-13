package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/RhykerWells/RK/backend/internal/config"
	"github.com/RhykerWells/RK/backend/internal/logger"
	"goji.io/v3"
)

var log *slog.Logger
var Multiplexer *goji.Mux

func Start() {
	log = logger.With("m", "server")

	mux := goji.NewMux()
	Multiplexer = mux

	address := fmt.Sprintf("%s:%d",
		config.AppConfig.Server.BindAddress,
		config.AppConfig.Server.Port,
	)

	log.Info(fmt.Sprintf("Webserver starting on %s", address))
	err := http.ListenAndServe(address, Multiplexer)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to start webserver: %v", err))
		os.Exit(1)
	}
}
