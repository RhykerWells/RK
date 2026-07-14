package server

import (
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"goji.io/v3/pat"
)

func (s *Server) registerRoutes() {
	s.Multiplexer.HandleFunc(pat.Get("/health"), handlers.Health)
}
