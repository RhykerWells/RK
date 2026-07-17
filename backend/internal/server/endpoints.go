package server

var ApiVersion = "v1"

var (
	EndpointAPI = "/api/" + ApiVersion

	EndpointHealth = EndpointAPI + "/health"

	// Authentication
	EndpointAuthLogin    = EndpointAPI + "/login"
	EndpointAuthCallback = EndpointAPI + "/callback"
	EndpointAuthLogout   = EndpointAPI + "/logout"

	// Users
	EndpointUsers = EndpointAPI + "/users"
	EndpointUser  = EndpointUsers + "/:user_id"
	EndpointMe    = EndpointUsers + "/@me"
	EndpointMeAPI = EndpointMe + "/api"

	// Portals
	EndpointPortals = EndpointAPI + "/portals"
	EndpointPortal  = EndpointPortals + "/:portal_id"

	// Portal Roles
	EndpointPortalRoles = EndpointPortal + "/roles"
	EndpointPortalRole  = EndpointPortalRoles + "/:role_id"

	// Portal Members
	EndpointPortalMembers    = EndpointPortal + "/members"
	EndpointPortalMember     = EndpointPortalMembers + "/:user_id"
	EndpointPortalMemberSelf = EndpointPortalMembers + "/@me"

	// Folders
	EndpointPortalFolders = EndpointPortal + "/folders"
	EndpointPortalFolder  = EndpointPortalFolders + "/:folder_id"

	// Documents
	EndpointFolderDocuments = EndpointPortal + "/documents"
	EndpointFolderDocument  = EndpointFolderDocuments + "/:document_id"
)
