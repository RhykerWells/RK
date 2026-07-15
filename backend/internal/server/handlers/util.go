package handlers

import (
	"context"
	"strconv"

	"github.com/RhykerWells/RK/backend/internal/app/users"
	"github.com/RhykerWells/RK/backend/internal/database/models"
	. "github.com/RhykerWells/RK/backend/internal/server/errors"
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
