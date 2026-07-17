package folders

import (
	"context"
	"time"

	"github.com/RhykerWells/RK/backend/internal/database/models"
	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
)

func GetPortalFolderByID(ctx context.Context, portalModel *models.Portal, folderID int64) (*models.Folder, error) {
	return models.Folders(models.FolderWhere.PortalID.EQ(portalModel.ID), models.FolderWhere.ID.EQ(folderID)).One(ctx, boil.GetContextDB())
}

func FolderCreate(ctx context.Context, portalID int64, req *CreateFolderRequest, userID int64) (*models.Folder, error) {
	newFolder := &models.Folder{
		PortalID:  portalID,
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if req.ParentFolderID != nil {
		newFolder.ParentFolderID = null.Int64{Int64: *req.ParentFolderID, Valid: true}
	}

	err := newFolder.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, err
	}

	if err = newFolder.Reload(ctx, boil.GetContextDB()); err != nil {
		return nil, err
	}

	return newFolder, nil
}

func FolderUpdate(ctx context.Context, folder *models.Folder, req *UpdateFolderRequest, userID int64) (*models.Folder, error) {
	updated := *folder

	if req.Name != nil {
		updated.Name = *req.Name
	}

	if req.ParentFolderID != nil {
		updated.ParentFolderID = *req.ParentFolderID
	}

	updated.UpdatedAt = time.Now()

	_, err := updated.Update(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, err
	}

	if err = updated.Reload(ctx, boil.GetContextDB()); err != nil {
		return nil, err
	}

	return &updated, nil
}

func FolderDelete(ctx context.Context, folder *models.Folder) error {
	_, err := folder.Delete(ctx, boil.GetContextDB())
	return err
}
