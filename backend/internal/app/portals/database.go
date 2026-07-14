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