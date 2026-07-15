package portals

import (
	"context"
	"time"

	"github.com/RhykerWells/RK/backend/internal/database/models"
	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
)

func GetPortalByID(ctx context.Context, id int64) (*models.Portal, error) {
	p, e := models.FindPortal(ctx, boil.GetContextDB(), id)

	return p, e
}

func PortalCreate(ctx context.Context, portal *CreatePortalRequest) (*models.Portal, error) {
	newPortal := &models.Portal{
		Name:      portal.Name,
		Domain:    portal.Domain,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

func PortalDelete(ctx context.Context, portal *models.Portal) error {
	_, err := portal.Delete(ctx, boil.GetContextDB())

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
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
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

	role.UpdatedAt = time.Now()

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
	newMembership := &models.PortalMembership{
		PortalID:  portalModel.ID,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

func MemberUpdate(ctx context.Context, memberModel *models.PortalMembership, update *UpdatePortalMemberRequest) (*models.PortalMembership, error) {
	member := *memberModel

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

	// Update membership timestamp
	member.UpdatedAt = time.Now()
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
