package portals

import (
	"context"
	"time"

	"github.com/RhykerWells/RK/backend/internal/database/models"
	"github.com/aarondl/sqlboiler/v4/boil"
)

func GetPortalByID(ctx context.Context, id int64) (*models.Portal, error) {
	p, e := models.FindPortal(ctx, boil.GetContextDB(), id)

	return p, e
}

func PortalCreate(ctx context.Context, portal *CreatePortalRequest) (*models.Portal, error) {
	newPortal := & models.Portal{
		Name: portal.Name,
		Domain: portal.Domain,
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

func GetPortalMemberByID(ctx context.Context, portalModel *models.Portal, userID int64) (*models.PortalMembership, error) {
	return models.PortalMemberships(models.PortalMembershipWhere.PortalID.EQ(portalModel.ID), models.PortalMembershipWhere.UserID.EQ(userID)).One(ctx, boil.GetContextDB())
}

func PortalMemberCreate(ctx context.Context, portalModel *models.Portal, userID int64) (*models.PortalMembership, error) {
	newMembership := &models.PortalMembership{
		PortalID: portalModel.ID,
		UserID: userID,
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

func MemberUpdate(ctx context.Context, portalModel *models.Portal, userID int64, update *UpdatePortalMemberRequest) (*models.PortalMembership, error) {
	member, _ := GetPortalMemberByID(ctx, portalModel, userID)

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

	return member, nil
}

func MemberDelete(ctx context.Context, portalModel *models.Portal, userID int64) error {
	member, _ := GetPortalMemberByID(ctx, portalModel, userID)

	_, err := member.Delete(ctx, boil.GetContextDB())

	return err
}