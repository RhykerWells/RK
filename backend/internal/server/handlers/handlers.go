package handlers

import "log/slog"

var log *slog.Logger

func SetLogger(logger *slog.Logger) {
	log = logger
}
