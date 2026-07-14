package users

import (
	"time"

	"github.com/aarondl/null/v8"
)

type User struct {
	ID              int64
	AuthType        string
	Username        string
	DisplayName     string
	Email           *string
	DiscordID       *string
	AvatarURL       *string
	IsActive        bool
	LastLoginAt     *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	IsAdministrator bool
}

type CreateUserRequest struct {
	Username        string      `json:"username"`
	DiscordID       null.String `json:"discord_id,omitempty"`
	DisplayName     string      `json:"display_name"`
	Email           null.String `json:"email,omitempty"`
	AvatarURL       null.String `json:"avatar_url,omitempty"`
	IsAdministrator bool        `json:"is_administrator"`
}

type UserResponse struct {
	ID              int64       `json:"id"`
	AuthType        string      `json:"auth_type"`
	Username        string      `json:"username"`
	DiscordID       null.String `json:"discord_id,omitempty"`
	DisplayName     string      `json:"display_name"`
	Email           null.String `json:"email,omitempty"`
	AvatarURL       null.String `json:"avatar_url,omitempty"`
	LastLoginAt     *time.Time  `json:"last_login_at,omitempty"`
	IsAdministrator bool        `json:"is_administrator"`
}
