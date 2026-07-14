package server

var ApiVersion = "v1"

var (
	EndpointAPI = "/api/" + ApiVersion

	EndpointHealth = EndpointAPI + "/health"

	// Users
	EndpointUsers      = EndpointAPI + "/users"
	EndpointUserCreate = EndpointUsers + "/create"
	EndpointUser       = EndpointUsers + "/:id"
	EndpointUserDelete = EndpointUser + "/delete"
)
