package server

import (
	"github.com/RhykerWells/RK/backend/internal/permissions"
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
	api.Handle(pat.Post(EndpointPortals), middleware.WithPermissionsMW(handlers.PortalCreate, permissions.PermissionPortalsManage))
	portalRequiredMux.Handle(pat.Delete(EndpointPortal), middleware.WithPermissionsMW(handlers.PortalDelete, permissions.PermissionPortalsManage))

	// Portal Role Retrieval Endpoints
	portalRequiredMux.HandleFunc(pat.Get(EndpointPortalRoles), handlers.PortalRoles)
	portalRequiredMux.HandleFunc(pat.Get(EndpointPortalRole), handlers.PortalRole)

	// Portal Role Management Endpoints
	portalRequiredMux.Handle(pat.Post(EndpointPortalRoles), middleware.WithPermissionsMW(handlers.PortalRoleCreate, permissions.PermissionPortalManageRoles))
	portalRequiredMux.Handle(pat.Patch(EndpointPortalRoles), middleware.WithPermissionsMW(handlers.PortalRoleUpdate, permissions.PermissionPortalManageRoles))
	portalRequiredMux.Handle(pat.Delete(EndpointPortalRoles), middleware.WithPermissionsMW(handlers.PortalRoleDelete, permissions.PermissionPortalManageRoles))

	// Portal Member Retrieval Endpoints
	portalRequiredMux.HandleFunc(pat.Get(EndpointPortalMembers), handlers.PortalMembers)
	portalRequiredMux.HandleFunc(pat.Get(EndpointPortalMember), handlers.PortalMember)

	// Portal Member Management Endpoints
	portalRequiredMux.Handle(pat.Post(EndpointPortalMember), middleware.WithPermissionsMW(handlers.PortalMemberCreate, permissions.PermissionPortalManageMembers))
	portalRequiredMux.Handle(pat.Patch(EndpointPortalMember), middleware.WithPermissionsMW(handlers.PortalMemberUpdate, permissions.PermissionPortalManageMembers))
	portalRequiredMux.Handle(pat.Delete(EndpointPortalMember), middleware.WithPermissionsMW(handlers.PortalMemberDelete, permissions.PermissionPortalManageMembers))

	api.Handle(pat.New("/*"), portalRequiredMux)
}
