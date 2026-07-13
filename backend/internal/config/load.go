package config

import (
	"context"
	"log/slog"
	"os"

	"github.com/RhykerWells/RK/backend/internal/log"
	"github.com/sethvargo/go-envconfig"
)

func Load() *Config {
	var log = log.With("m", "config")

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

	log.Info(
		"Configuration loaded successfully",
	)

	return &cfg
}
