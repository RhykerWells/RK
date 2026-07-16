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
	handlers.InitDiscordOauth()

	// Health
	s.Multiplexer.HandleFunc(pat.Get(EndpointHealth), handlers.Health)

	// Frontend auth
	s.Multiplexer.HandleFunc(pat.Get(EndpointAuthLogin), handlers.Login)
	s.Multiplexer.HandleFunc(pat.Get(EndpointAuthCallback), handlers.Callback)
	s.Multiplexer.HandleFunc(pat.Get(EndpointAuthLogout), handlers.Logout)
	s.Multiplexer.Handle(pat.Get(EndpointMe), middleware.WithAuthMW(http.HandlerFunc(handlers.Me)))

	// API token generation
	s.Multiplexer.Handle(pat.Post(EndpointMeAPI), middleware.WithAuthMW(http.HandlerFunc(handlers.IssueAPIToken)))
	s.Multiplexer.Handle(pat.Delete(EndpointMeAPI), middleware.WithAuthMW(http.HandlerFunc(handlers.RevokeAPIToken)))

	// Users
	s.Multiplexer.Handle(pat.Get(EndpointUser), middleware.WithAuthMW(http.HandlerFunc(handlers.User)))
	s.Multiplexer.HandleFunc(pat.Post(EndpointUserCreate), handlers.UserCreate)
	s.Multiplexer.Handle(pat.Delete(EndpointUserDelete), middleware.WithAuthMW(http.HandlerFunc(handlers.UserDelete)))

	// Portals
	s.Multiplexer.Handle(pat.Get(EndpointPortal), middleware.WithAuthMW(middleware.WithPortalMW(http.HandlerFunc(handlers.Portal))))
	s.Multiplexer.Handle(pat.Get(EndpointPortalCreate), middleware.WithAuthMW(http.HandlerFunc(handlers.PortalCreate)))
	s.Multiplexer.Handle(pat.Delete(EndpointPortalDelete), middleware.WithAuthMW(middleware.WithPortalMW(http.HandlerFunc(handlers.PortalDelete))))

	// Portal Roles
	s.Multiplexer.Handle(pat.Get(EndpointPortalRoles), middleware.WithAuthMW(middleware.WithPortalMW(http.HandlerFunc(handlers.PortalRoles))))
	s.Multiplexer.Handle(pat.Get(EndpointPortalRole), middleware.WithAuthMW(middleware.WithPortalMW(http.HandlerFunc(handlers.PortalRole))))
	s.Multiplexer.Handle(pat.Post(EndpointPortalRoles), middleware.WithAuthMW(middleware.WithPortalMW(http.HandlerFunc(handlers.PortalRoleCreate))))
	s.Multiplexer.Handle(pat.Patch(EndpointPortalRole), middleware.WithAuthMW(middleware.WithPortalMW(http.HandlerFunc(handlers.PortalRoleUpdate))))
	s.Multiplexer.Handle(pat.Delete(EndpointPortalRole), middleware.WithAuthMW(middleware.WithPortalMW(http.HandlerFunc(handlers.PortalRoleDelete))))

	// Portal Members
	s.Multiplexer.Handle(pat.Get(EndpointPortalMembers), middleware.WithAuthMW(middleware.WithPortalMW(http.HandlerFunc(handlers.PortalMembers))))
	s.Multiplexer.Handle(pat.Get(EndpointPortalMember), middleware.WithAuthMW(middleware.WithPortalMW(http.HandlerFunc(handlers.PortalMember))))
	s.Multiplexer.Handle(pat.Post(EndpointPortalMember), middleware.WithAuthMW(middleware.WithPortalMW(http.HandlerFunc(handlers.PortalMemberCreate))))
	s.Multiplexer.Handle(pat.Patch(EndpointPortalMember), middleware.WithAuthMW(middleware.WithPortalMW(http.HandlerFunc(handlers.PortalMemberUpdate))))
	s.Multiplexer.Handle(pat.Delete(EndpointPortalMember), middleware.WithAuthMW(middleware.WithPortalMW(http.HandlerFunc(handlers.PortalMemberDelete))))
}
