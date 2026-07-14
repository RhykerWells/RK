package handlers

import (
	"context"
	"strconv"

	"github.com/RhykerWells/RK/backend/internal/app/portals"
	"github.com/RhykerWells/RK/backend/internal/app/users"
	"github.com/RhykerWells/RK/backend/internal/database/models"
)

func getUser(ctx context.Context, requestType string, id string) (*models.User, error) {
	var (
		user *models.User
		err  error
	)

	switch requestType {
	case "id":
		id, convErr := strconv.ParseInt(id, 10, 64)
		if convErr != nil {
			return nil, ErrInvalidUserID
		}
		user, err = users.GetUserByID(ctx, id)
	case "discord":
		user, err = users.GetUserByDiscordID(ctx, id)
	default:
		return nil, ErrInvalidMissingRequestType
	}

	return user, err
}

func getPortal(ctx context.Context, idStr string) (*models.Portal, error) {
	id, convErr := strconv.ParseInt(idStr, 10, 64)
	if convErr != nil {
		return nil, ErrInvalidPortalID
	}

	return portals.GetPortalByID(ctx, id)
}

func getPortalRoles(portalModel *models.Portal) ([]portals.PortalRoleResponse) {
	roles := make([]portals.PortalRoleResponse, 0, len(portalModel.GetPortalRoles()))

	for _, role := range portalModel.GetPortalRoles() {
		roles = append(roles, portals.PortalRoleResponse{
			ID:            role.ID,
			Name:          role.Name,
			Description:   role.Description,
			Colour:        role.Colour,
			Position:      role.Position,
			DiscordRoleID: role.DiscordRoleID,
			CreatedAt:     role.CreatedAt,
			UpdatedAt:     role.UpdatedAt,
		})
	}

	return roles
}