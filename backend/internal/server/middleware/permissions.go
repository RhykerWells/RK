package middleware

import (
	"github.com/RhykerWells/RK/backend/internal/app/portals"
	"github.com/RhykerWells/RK/backend/internal/permissions"
)

func hasRequiredPermission(roles []portals.PortalRoleResponse, required []permissions.Permission) bool {
	for _, role := range roles {
		for _, perm := range role.Permissions {
			if perm == string(permissions.PermissionAdministrator) {
				return true
			}

			for _, requiredPerm := range required {
				if perm == string(requiredPerm) {
					return true
				}
			}
		}
	}

	return false
}
