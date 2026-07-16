package middleware

import (
	"slices"

	"github.com/RhykerWells/RK/backend/internal/app/portals"
	"github.com/RhykerWells/RK/backend/internal/permissions"
)

func hasRequiredPermission(roles []portals.PortalRoleResponse, required []permissions.Permission) bool {
	for _, role := range roles {
		if slices.Contains(role.Permissions, permissions.PermissionName(permissions.PermissionAdministrator)) {
			return true
		}

		for _, perm := range role.Permissions {
			if slices.Contains(required, permissions.Permission(perm)) {
				return true
			}
		}
	}

	return false
}
