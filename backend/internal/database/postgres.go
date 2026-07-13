package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/RhykerWells/RK/backend/internal/config"
	"github.com/RhykerWells/RK/backend/internal/database/schemas"
	"github.com/RhykerWells/RK/backend/internal/log"
	"github.com/aarondl/sqlboiler/v4/boil"
	_ "github.com/lib/pq"
)

func Connect(config *config.Config) error {
	var log = log.With("m", "config")

	log.Info(
		"Connecting to Database",
	)

	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Name))
	if err != nil {
		log.Error(
			"Failed to create database handle",
			slog.Any("error", err),
		)
		os.Exit(1)
	}

	if err := DB.Ping(); err != nil {
		log.Error(
			"Failed to connect to database",
			slog.String(
				"database",
				fmt.Sprintf(
					"%s@%s:%d/%s",
					config.Database.User,
					config.Database.Host,
					config.Database.Port,
					config.Database.Name,
				),
			),
			slog.Any("error", err),
		)
		os.Exit(1)
	}

	boil.SetDB(DB)

	log.Info(
		"Database connected successfully",
	)

	initSchemas(log)

	return nil
}

func initSchemas(log *slog.Logger) {
	log.Info(
		"Initialising database schemas",
	)

	for _, schemaGroup := range schemas.AllSchemas() {
		for _, schema := range schemaGroup {
			_, err := DB.Exec(schema.SQL)
			if err != nil {
				log.Error(
					"Failed initialising postgres db schema",
					slog.String("schema", schema.Name),
					slog.Any("error", err),
				)
				os.Exit(1)
				return
			}

			log.Info("Successfully initialised postgres db schema", slog.String("schema", schema.Name))
		}

	}

	log.Info("All database schemas initialised successfully")
}
