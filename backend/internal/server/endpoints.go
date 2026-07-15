package server

var ApiVersion = "v1"

var (
	EndpointAPI = "/api/" + ApiVersion

	EndpointHealth = EndpointAPI + "/health"

	// Frontend auth
	EndpointAuthLogin    = EndpointAPI + "/login"
	EndpointAuthCallback = EndpointAPI + "/callback"
	EndpointAuthLogout   = EndpointAPI + "/logout"

	// Users
	EndpointUsers      = EndpointAPI + "/users"
	EndpointUserCreate = EndpointUsers + "/create"
	EndpointUser       = EndpointUsers + "/:user_id"
	EndpointUserDelete = EndpointUser + "/delete"

	// Portals
	EndpointPortals      = EndpointAPI + "/portals"
	EndpointPortalCreate = EndpointPortals + "/create"
	EndpointPortal       = EndpointPortals + "/:portal_id"
	EndpointPortalDelete = EndpointPortal + "/delete"

	// Portal Roles
	EndpointPortalRoles = EndpointPortal + "/roles"
	EndpointPortalRole  = EndpointPortalRoles + "/:role_id"

	// Portal Members
	EndpointPortalMembers = EndpointPortal + "/members"
	EndpointPortalMember  = EndpointPortalMembers + "/:user_id"
)
