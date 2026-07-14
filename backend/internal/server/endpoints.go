package server

var ApiVersion = "v1"

var (
	EndpointAPI = "/api/" + ApiVersion

	EndpointHealth = EndpointAPI + "/health"

	// Users
	EndpointUsers      = EndpointAPI + "/users"
	EndpointUserCreate = EndpointUsers + "/create"
	EndpointUser       = EndpointUsers + "/:user_id"
	EndpointUserDelete = EndpointUser + "/delete"

	// Portals
	EndpointPortals = EndpointAPI + "/portals"
	EndpointPortalCreate = EndpointPortals + "/create"
	EndpointPortal = EndpointPortals + "/:portal_id"
	EndpointPortalDelete = EndpointPortal + "/delete"
)
