package portals

import (
	"time"

	"github.com/RhykerWells/RK/backend/internal/app/users"
	"github.com/aarondl/null/v8"
)

type Portal struct {
	ID        int64
	Name      string
	Domain    string
	Roles     []PortalRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreatePortalRequest struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

type PortalResponse struct {
	ID        int64                `json:"id"`
	Name      string               `json:"name"`
	Domain    string               `json:"domain"`
	Roles     []PortalRoleResponse `json:"roles,omitempty"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

type PortalRole struct {
	ID            int64
	PortalID      int64
	Name          string
	Description   null.String
	Colour        null.String
	Position      int64
	DiscordRoleID null.String
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CreatePortalRoleRequest struct {
	Name          string      `json:"name"`
	Description   null.String `json:"description,omitempty"`
	DiscordRoleID null.String `json:"discord_role_id,omitempty"`
}

type UpdatePortalRoleRequest struct {
	Name          *string `json:"name"`
	Description   *string `json:"description,omitempty"`
	Colour        *string `json:"colour,omitempty"`
	Position      *int64  `json:"position,omitempty"`
	DiscordRoleID *string `json:"discord_role_id,omitempty"`
}

type PortalRoleResponse struct {
	ID            int64       `json:"id"`
	Name          string      `json:"name"`
	Description   null.String `json:"description,omitempty"`
	Colour        null.String `json:"colour,omitempty"`
	Position      int64       `json:"position"`
	DiscordRoleID null.String `json:"discord_role_id,omitempty"`
	Permissions   []string    `json:"permissions,omitempty"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

type PortalMember struct {
	ID        int64
	PortalID  int64
	User      users.User
	Roles     []int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdatePortalMemberRequest struct {
	Roles []int64 `json:"roles"`
}

type PortalMemberResponse struct {
	Roles     []PortalRoleResponse `json:"roles,omitempty"`
	User      users.UserResponse   `json:"user"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}
