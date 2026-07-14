package server

import (
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"goji.io/v3/pat"
)

// RegisterRoutes registers all API routes on the server's multiplexer.
// Handlers are assumed to be implemented separately
// These are only route definitions.
func (s *Server) registerRoutes() {
	// Health
	s.Multiplexer.HandleFunc(pat.Get(EndpointHealth), handlers.Health)

	// Users
	s.Multiplexer.HandleFunc(pat.Get(EndpointUser), handlers.User)
	s.Multiplexer.HandleFunc(pat.Post(EndpointUserCreate), handlers.UserCreate)
	s.Multiplexer.HandleFunc(pat.Delete(EndpointUserDelete), handlers.UserDelete)

	// Portals
	s.Multiplexer.HandleFunc(pat.Get(EndpointPortal), handlers.Portal)
	s.Multiplexer.HandleFunc(pat.Get(EndpointPortalCreate), handlers.PortalCreate)
	s.Multiplexer.HandleFunc(pat.Delete(EndpointPortalDelete), handlers.PortalDelete)
}
