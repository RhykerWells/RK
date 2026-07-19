package users

import "github.com/RhykerWells/RK/backend/internal/database/models"

func UsersModelToResponse(models []*models.User) []UserResponse {
	responses := make([]UserResponse, len(models))

	for i, model := range models {
		responses[i] = UserModelToResponse(model)
	}

	return responses
}

func UserModelToResponse(model *models.User) UserResponse {
	user := UserResponse{
		ID:              model.ID,
		AuthType:        model.AuthType,
		DiscordID:       model.DiscordID,
		Username:        model.Username,
		DisplayName:     model.DisplayName,
		Email:           model.Email,
		IsAdministrator: model.IsAdministrator,
	}

	if model.LastLoginAt.Valid {
		t := model.LastLoginAt.Time
		user.LastLoginAt = &t
	}

	return user
}
