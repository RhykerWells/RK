package portals

import (
	"context"
	"database/sql"
	"errors"

	"github.com/RhykerWells/RK/backend/internal/app/users"
	"github.com/RhykerWells/RK/backend/internal/database/models"
	. "github.com/RhykerWells/RK/backend/internal/server/errors"
	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
)

func GetPortalByID(ctx context.Context, id int64) (*models.Portal, error) {
	p, e := models.FindPortal(ctx, boil.GetContextDB(), id)

	return p, e
}

func PortalCreate(ctx context.Context, portal *CreatePortalRequest) (*models.Portal, error) {
	newPortal := &models.Portal{
		Name:   portal.Name,
		Domain: portal.Domain,
	}

	err := newPortal.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, err
	}

	if err = newPortal.Reload(ctx, boil.GetContextDB()); err != nil {
		return nil, err
	}

	return newPortal, nil
}

func PortalUpdate(ctx context.Context, portalModel *models.Portal, req *UpdatePortalRequest) (*models.Portal, error) {
	updated := *portalModel

	if req.Name != nil {
		updated.Name = *req.Name
	}

	if req.Domain != nil {
		updated.Domain = *req.Domain
	}

	_, err := updated.Update(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, err
	}

	if err = updated.Reload(ctx, boil.GetContextDB()); err != nil {
		return nil, err
	}

	return &updated, nil
}

func PortalDelete(ctx context.Context, portalModel *models.Portal) error {
	_, err := portalModel.Delete(ctx, boil.GetContextDB())

	return err
}

func GetPortalRoleByID(ctx context.Context, portalModel *models.Portal, roleID int64) (*models.PortalRole, error) {
	return models.PortalRoles(models.PortalRoleWhere.PortalID.EQ(portalModel.ID), models.PortalRoleWhere.ID.EQ(roleID)).One(ctx, boil.GetContextDB())
}

func PortalRoleCreate(ctx context.Context, portalModel *models.Portal, role *CreatePortalRoleRequest) (*models.PortalRole, error) {
	newRole := &models.PortalRole{
		PortalID:      portalModel.ID,
		Name:          role.Name,
		Description:   role.Description,
		DiscordRoleID: role.DiscordRoleID,
	}

	err := newRole.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, err
	}

	if err = newRole.Reload(ctx, boil.GetContextDB()); err != nil {
		return nil, err
	}

	return newRole, nil
}

func PortalRoleUpdate(ctx context.Context, roleModel *models.PortalRole, update *UpdatePortalRoleRequest) (*models.PortalRole, error) {
	role := *roleModel

	if update.Name != nil {
		role.Name = *update.Name
	}

	if update.Description != nil {
		role.Description = null.StringFromPtr(update.Description)
	}

	if update.Colour != nil {
		role.Colour = null.StringFromPtr(update.Colour)
	}

	if update.Position != nil {
		role.Position = *update.Position
	}

	if update.DiscordRoleID != nil {
		role.DiscordRoleID = null.StringFromPtr(update.DiscordRoleID)
	}

	if _, err := role.Update(ctx, boil.GetContextDB(), boil.Infer()); err != nil {
		return nil, err
	}

	return &role, nil
}

func PortalRoleDelete(ctx context.Context, roleModel *models.PortalRole) error {
	_, err := roleModel.Delete(ctx, boil.GetContextDB())

	return err
}

func GetPortalMemberByID(ctx context.Context, portalModel *models.Portal, userID int64) (*models.PortalMembership, error) {
	return models.PortalMemberships(models.PortalMembershipWhere.PortalID.EQ(portalModel.ID), models.PortalMembershipWhere.UserID.EQ(userID)).One(ctx, boil.GetContextDB())
}

func PortalMemberCreate(ctx context.Context, portalModel *models.Portal, userID int64) (*models.PortalMembership, error) {
	if _, err := users.GetUserByID(ctx, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if _, err := GetPortalMemberByID(ctx, portalModel, userID); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	} else {
		return nil, ErrMemberAlreadyExists
	}

	newMembership := &models.PortalMembership{
		PortalID: portalModel.ID,
		UserID:   userID,
	}

	err := newMembership.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, err
	}

	if err = newMembership.Reload(ctx, boil.GetContextDB()); err != nil {
		return nil, err
	}

	return newMembership, nil
}

func PortalMemberUpdate(ctx context.Context, memberModel *models.PortalMembership, update *UpdatePortalMemberRequest) (*models.PortalMembership, error) {
	member := *memberModel

	// Validate the requested roles belong to the membership portal.
	for _, roleID := range update.Roles {
		if _, err := GetPortalRoleByID(ctx, &models.Portal{ID: member.PortalID}, roleID); err != nil {
			return nil, err
		}
	}

	// Remove existing roles
	_, err := models.PortalMembershipRoles(models.PortalMembershipRoleWhere.PortalMembershipID.EQ(member.ID)).DeleteAll(ctx, boil.GetContextDB())
	if err != nil {
		return nil, err
	}

	// Add new roles
	for _, roleID := range update.Roles {
		role := &models.PortalMembershipRole{
			PortalMembershipID: member.ID,
			PortalRoleID:       roleID,
		}

		if err := role.Insert(ctx, boil.GetContextDB(), boil.Infer()); err != nil {
			return nil, err
		}

		if err := role.Reload(ctx, boil.GetContextDB()); err != nil {
			return nil, err
		}
	}

	if _, err := member.Update(ctx, boil.GetContextDB(), boil.Infer()); err != nil {
		return nil, err
	}

	// Reload relationships so response has the new roles
	if err := member.Reload(ctx, boil.GetContextDB()); err != nil {
		return nil, err
	}

	return &member, nil
}

func MemberDelete(ctx context.Context, memberModel *models.PortalMembership) error {
	_, err := memberModel.Delete(ctx, boil.GetContextDB())

	return err
}
