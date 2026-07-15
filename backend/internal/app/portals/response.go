package portals

import (
	"github.com/RhykerWells/RK/backend/internal/app/users"
	"github.com/RhykerWells/RK/backend/internal/database/models"
)

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
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}

func PortalMembersFromModel(model *models.Portal) ([]PortalMemberResponse) {
	members := make([]PortalMemberResponse, 0, len(model.GetPortalMemberships()))

	for _, member := range model.GetPortalMemberships() {

		members = append(members, PortalMemberFromModel(member))
	}

	return members
}

func PortalMemberFromModel(model *models.PortalMembership) PortalMemberResponse {
	roleIDs := make([]int64, 0, len(model.GetPortalMembershipRoles()))

	for _, role := range model.GetPortalMembershipRoles() {
		roleIDs = append(roleIDs, role.PortalRoleID)
	}

	user := model.GetUser()

	return PortalMemberResponse{
		Roles: roleIDs,
		User: users.UserModelToResponse(user),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}