package users

import (
	"context"

	"github.com/RhykerWells/RK/backend/internal/database/models"
	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
)

type Authtype string

const (
	AuthTypeDiscord Authtype = "discord"
	AuthTypeLocal   Authtype = "local"
)

func GetUsers(ctx context.Context) ([]*models.User, error) {
	users, err := models.Users().All(ctx, boil.GetContextDB())

	return users, err
}

func GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	u, e := models.FindUser(ctx, boil.GetContextDB(), id)

	return u, e
}

func GetUserByDiscordID(ctx context.Context, discordID string) (*models.User, error) {
	return models.Users(models.UserWhere.DiscordID.EQ(null.StringFrom(discordID))).One(ctx, boil.GetContextDB())
}

func UserCreate(ctx context.Context, user *CreateUserRequest, authType Authtype) (*models.User, error) {
	newUser := &models.User{
		AuthType:        string(authType),
		DiscordID:       user.DiscordID,
		Username:        user.Username,
		DisplayName:     user.DisplayName,
		Email:           user.Email,
		AvatarURL:       user.AvatarURL,
		IsAdministrator: false,
	}

	err := newUser.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, err
	}

	if err = newUser.Reload(ctx, boil.GetContextDB()); err != nil {
		return nil, err
	}

	return newUser, nil
}

func UserDelete(ctx context.Context, user *models.User) error {
	_, err := user.Delete(ctx, boil.GetContextDB())

	return err
}
