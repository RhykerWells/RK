package portals

import (
	"sort"

	"github.com/RhykerWells/RK/backend/internal/app/users"
	"github.com/RhykerWells/RK/backend/internal/database/models"
)

func PortalsModelToResponse(models []*models.Portal) []PortalResponse {
	responses := make([]PortalResponse, 0, len(models))

	for _, model := range models {
		responses = append(responses, PortalModelToResponse(model))
	}

	return responses
}

func PortalModelToResponse(model *models.Portal) PortalResponse {
	return PortalResponse{
		ID:        model.ID,
		Name:      model.Name,
		Domain:    model.Domain,
		Roles:     PortalRolesFromModel(model),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func PortalRolesFromModel(model *models.Portal) []PortalRoleResponse {
	roles := make([]PortalRoleResponse, 0, len(model.GetPortalRoles()))

	for _, role := range model.GetPortalRoles() {
		roles = append(roles, PortalRoleFromModel(role))
	}

	return roles
}

func PortalRoleFromModel(model *models.PortalRole) PortalRoleResponse {
	return PortalRoleResponse{
		ID:            model.ID,
		Name:          model.Name,
		Description:   model.Description,
		Colour:        model.Colour,
		Position:      model.Position,
		DiscordRoleID: model.DiscordRoleID,
		Permissions:   PortalRolePermissionsFromModel(model),
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}

func PortalRolePermissionsFromModel(model *models.PortalRole) []string {
	permissionSet := make(map[string]struct{})
	for _, rolePermission := range model.GetPortalRolePermissions() {
		permissionSet[rolePermission.PermissionKey] = struct{}{}
	}

	permissions := make([]string, 0, len(permissionSet))
	for permission := range permissionSet {
		permissions = append(permissions, permission)
	}
	sort.Strings(permissions)

	return permissions
}

func PortalMembersFromModel(model *models.Portal) []PortalMemberResponse {
	members := make([]PortalMemberResponse, 0, len(model.GetPortalMemberships()))

	for _, member := range model.GetPortalMemberships() {
		members = append(members, PortalMemberFromModel(member))
	}

	return members
}

func PortalMemberFromModel(model *models.PortalMembership) PortalMemberResponse {
	roles := make([]PortalRoleResponse, 0, len(model.GetPortalMembershipRoles()))

	for _, membershipRole := range model.GetPortalMembershipRoles() {
		role := membershipRole.GetPortalRole()
		if role == nil {
			continue
		}
		roles = append(roles, PortalRoleFromModel(role))
	}

	user := model.GetUser()

	return PortalMemberResponse{
		Roles:     roles,
		User:      users.UserModelToResponse(user),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
