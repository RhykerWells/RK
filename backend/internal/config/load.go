package config

import (
	"context"
	"log/slog"
	"os"

	"github.com/RhykerWells/RK/backend/internal/logger"
	"github.com/sethvargo/go-envconfig"
)

var AppConfig *Config
var log *slog.Logger

func Load() *Config {
	log = logger.With("m", "config")

	log.Info(
		"Loading Configuration",
	)

	var cfg Config
	err := envconfig.Process(context.Background(), &cfg)

	if err != nil {
		log.Error(
			"Failed to load configuration",
			slog.Any("error", err),
		)
		os.Exit(1)
	}

	AppConfig = &cfg

	log.Info(
		"Configuration loaded successfully",
	)

	return &cfg
}
