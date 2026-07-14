package server

import (
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"goji.io/v3/pat"
)

func registerRoutes() {
	Multiplexer.HandleFunc(pat.Get("/health"), handlers.Health)
}
