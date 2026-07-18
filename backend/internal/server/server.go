package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RhykerWells/RK/backend/internal/config"
	"github.com/RhykerWells/RK/backend/internal/logger"
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router chi.Router
	Logger *slog.Logger
}

func NewServer() *Server {
	s := &Server{
		Router: chi.NewRouter(),
		Logger: logger.With("m", "server"),
	}

	handlers.SetLogger(s.Logger)
	s.registerRoutes()

	return s
}

func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%d",
		config.AppConfig.Server.BindAddress,
		config.AppConfig.Server.Port,
	)

	httpServer := &http.Server{
		Addr:    address,
		Handler: s.Router,
	}

	errs := make(chan error, 1)
	// Start listening in the background
	go func() {
		s.Logger.Info("Starting weberver listen and serve",
			"addr", address,
		)

		if err := httpServer.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			errs <- err
		}
	}()

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errs:
		s.Logger.Error("Webserver failed to listen and serve",
			"error", err,
		)

		return err

	case <-quit:
		s.Logger.Info("Webserver shutting down")

		ctx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			return err
		}

		s.Logger.Info("Webserver stopped")
		return nil
	}
}
