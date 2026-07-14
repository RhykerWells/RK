package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/RhykerWells/RK/backend/internal/config"
	"github.com/RhykerWells/RK/backend/internal/logger"
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"goji.io/v3"
)

var log *slog.Logger
var Multiplexer *goji.Mux

func Start() {
	log = logger.With("m", "server")

	mux := goji.NewMux()
	Multiplexer = mux

	if config.AppConfig.Server.Debug {
		mux.Use(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				log.Info("Request received", "method", r.Method, "path", r.URL.Path)
				h.ServeHTTP(w, r)
			})
		})
	}

	handlers.SetLogger(log)
	registerRoutes()

	address := fmt.Sprintf(":%d",
		config.AppConfig.Server.Port,
	)

	log.Info("Webserver starting", "address", address)
	err := http.ListenAndServe(address, Multiplexer)
	if err != nil {
		log.Error("Failed to start webserver", "error", err)
		os.Exit(1)
	}
}
