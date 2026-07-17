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
	portalRequiredMux.Handle(pat.Get(EndpointPortal), middleware.WithPortalMembershipMW(handlers.Portal))

	// Portal Management Endpoints
	api.HandleFunc(pat.Post(EndpointPortals), handlers.PortalCreate)
	portalRequiredMux.Handle(pat.Delete(EndpointPortal), middleware.WithPermissionsMW(handlers.PortalDelete, permissions.PermissionPortalsManage))

	// Portal Role Retrieval Endpoints
	portalRequiredMux.Handle(pat.Get(EndpointPortalRoles), middleware.WithPortalMembershipMW(handlers.PortalRoles))
	portalRequiredMux.Handle(pat.Get(EndpointPortalRole), middleware.WithPortalMembershipMW(handlers.PortalRole))

	// Portal Role Management Endpoints
	portalRequiredMux.Handle(pat.Post(EndpointPortalRoles), middleware.WithPermissionsMW(handlers.PortalRoleCreate, permissions.PermissionPortalManageRoles))
	portalRequiredMux.Handle(pat.Patch(EndpointPortalRoles), middleware.WithPermissionsMW(handlers.PortalRoleUpdate, permissions.PermissionPortalManageRoles))
	portalRequiredMux.Handle(pat.Delete(EndpointPortalRoles), middleware.WithPermissionsMW(handlers.PortalRoleDelete, permissions.PermissionPortalManageRoles))

	// Portal Member Retrieval Endpoints
	portalRequiredMux.Handle(pat.Get(EndpointPortalMembers), middleware.WithPortalMembershipMW(handlers.PortalMembers))
	portalRequiredMux.Handle(pat.Get(EndpointPortalMember), middleware.WithPortalMembershipMW(handlers.PortalMember))

	// Portal Member Self-management Endpoint
	portalRequiredMux.HandleFunc(pat.Post(EndpointPortalMemberSelf), handlers.PortalMemberJoin)
	portalRequiredMux.HandleFunc(pat.Delete(EndpointPortalMemberSelf), handlers.PortalMemberLeave)

	// Portal Member Management Endpoints
	portalRequiredMux.Handle(pat.Post(EndpointPortalMember), middleware.WithPermissionsMW(handlers.PortalMemberCreate, permissions.PermissionPortalManageMembers))
	portalRequiredMux.Handle(pat.Patch(EndpointPortalMember), middleware.WithPermissionsMW(handlers.PortalMemberUpdate, permissions.PermissionPortalManageMembers, permissions.PermissionPortalManageMemberRoles))
	portalRequiredMux.Handle(pat.Delete(EndpointPortalMember), middleware.WithPermissionsMW(handlers.PortalMemberDelete, permissions.PermissionPortalManageMembers))

	api.Handle(pat.New("/*"), portalRequiredMux)
}
