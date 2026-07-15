package server

import (
	"net/http"

	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
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
	s.Multiplexer.Handle(pat.Get(EndpointPortal), middleware.WithPortalMW(http.HandlerFunc(handlers.Portal)))
	s.Multiplexer.HandleFunc(pat.Get(EndpointPortalCreate), handlers.PortalCreate)
	s.Multiplexer.Handle(pat.Delete(EndpointPortalDelete), middleware.WithPortalMW(http.HandlerFunc(handlers.PortalDelete)))

	// Portal Roles
	s.Multiplexer.Handle(pat.Get(EndpointPortalRoles), middleware.WithPortalMW(http.HandlerFunc(handlers.PortalRoles)))
	s.Multiplexer.Handle(pat.Get(EndpointPortalRole), middleware.WithPortalMW(http.HandlerFunc(handlers.PortalRole)))
	s.Multiplexer.Handle(pat.Post(EndpointPortalRoles), middleware.WithPortalMW(http.HandlerFunc(handlers.PortalRoleCreate)))
	s.Multiplexer.Handle(pat.Patch(EndpointPortalRole), middleware.WithPortalMW(http.HandlerFunc(handlers.PortalRoleUpdate)))
	s.Multiplexer.Handle(pat.Delete(EndpointPortalRole), middleware.WithPortalMW(http.HandlerFunc(handlers.PortalRoleDelete)))

	// Portal Members
	s.Multiplexer.Handle(pat.Get(EndpointPortalMembers), middleware.WithPortalMW(http.HandlerFunc(handlers.PortalMembers)))
	s.Multiplexer.Handle(pat.Get(EndpointPortalMember), middleware.WithPortalMW(http.HandlerFunc(handlers.PortalMember)))
	s.Multiplexer.Handle(pat.Post(EndpointPortalMember), middleware.WithPortalMW(http.HandlerFunc(handlers.PortalMemberCreate)))
	s.Multiplexer.Handle(pat.Patch(EndpointPortalMember), middleware.WithPortalMW(http.HandlerFunc(handlers.PortalMemberUpdate)))
	s.Multiplexer.Handle(pat.Delete(EndpointPortalMember), middleware.WithPortalMW(http.HandlerFunc(handlers.PortalMemberDelete)))
}
