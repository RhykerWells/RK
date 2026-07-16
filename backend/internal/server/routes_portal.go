package server

import (
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"goji.io/v3"
	"goji.io/v3/pat"
)

func registerPortalRoutes(api *goji.Mux) {
	portalRequiredMux := goji.SubMux()
	portalRequiredMux.Use(middleware.WithPortalMW)

	// Portal Retrieval Endpoints
	portalRequiredMux.HandleFunc(pat.Get(EndpointPortal), handlers.Portal)

	// Portal Management Endpoints
	api.HandleFunc(pat.Post(EndpointPortals), handlers.PortalCreate)

	portalRequiredMux.HandleFunc(pat.Delete(EndpointPortal), handlers.PortalDelete)

	// Portal Role Retrieval Endpoints
	portalRequiredMux.HandleFunc(pat.Get(EndpointPortalRoles), handlers.PortalRoles)
	portalRequiredMux.HandleFunc(pat.Get(EndpointPortalRole), handlers.PortalRole)

	// Portal Role Management Endpoints
	portalRequiredMux.HandleFunc(pat.Post(EndpointPortalRoles), handlers.PortalRoleCreate)
	portalRequiredMux.HandleFunc(pat.Patch(EndpointPortalRole), handlers.PortalRoleUpdate)
	portalRequiredMux.HandleFunc(pat.Delete(EndpointPortalRole), handlers.PortalRoleDelete)

	// Portal Member Retrieval Endpoints
	portalRequiredMux.HandleFunc(pat.Get(EndpointPortalMembers), handlers.PortalMembers)
	portalRequiredMux.HandleFunc(pat.Get(EndpointPortalMember), handlers.PortalMember)

	// Portal Member Management Endpoints
	portalRequiredMux.HandleFunc(pat.Post(EndpointPortalMember), handlers.PortalMemberCreate)
	portalRequiredMux.HandleFunc(pat.Patch(EndpointPortalMember), handlers.PortalMemberUpdate)
	portalRequiredMux.HandleFunc(pat.Delete(EndpointPortalMember), handlers.PortalMemberDelete)

	api.Handle(pat.New("/*"), portalRequiredMux)
}
