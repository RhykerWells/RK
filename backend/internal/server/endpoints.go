package server

var ApiVersion = "v1"

var (
	EndpointAPI = "/api/" + ApiVersion

	EndpointHealth = EndpointAPI + "/health"

	// Authentication
	EndpointAuthLogin    = EndpointAPI + "/login"
	EndpointAuthCallback = EndpointAPI + "/callback"
	EndpointAuthLogout   = EndpointAPI + "/logout"

	// Current user
	// Requires login through the dashboard
	EndpointMe    = EndpointAPI + "/me"
	EndpointMeAPI = EndpointMe + "/api"

	// Users
	EndpointUsers = EndpointAPI + "/users"
	EndpointUser  = EndpointUsers + "/:user_id"

	// Portals
	EndpointPortals = EndpointAPI + "/portals"
	EndpointPortal  = EndpointPortals + "/:portal_id"

	// Portal Roles
	EndpointPortalRoles = EndpointPortal + "/roles"
	EndpointPortalRole  = EndpointPortalRoles + "/:role_id"

	// Portal Members
	EndpointPortalMembers = EndpointPortal + "/members"
	EndpointPortalMember  = EndpointPortalMembers + "/:user_id"
)
