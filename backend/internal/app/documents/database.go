package documents

import (
	"context"

	"github.com/RhykerWells/RK/backend/internal/database/models"
	. "github.com/RhykerWells/RK/backend/internal/server/errors"
	"github.com/aarondl/sqlboiler/v4/boil"
)

func GetDocumentByID(ctx context.Context, id int64) (*models.Document, error) {
	d, e := models.FindDocument(ctx, boil.GetContextDB(), id)

	return d, e
}

func DocumentCreate(ctx context.Context, req *CreateDocumentRequest, portalModel *models.Portal, userModel *models.User) (*models.Document, error) {
	if req.Title == "" {
		return nil, ErrDocumentTitleRequired
	}

	if req.FolderID.Valid == req.OwnerID.Valid {
		return nil, ErrNewDocumentInvalidLocation
	}

	newDoc := &models.Document{
		PortalID:  portalModel.ID,
		Title:     req.Title,
		CreatedBy: userModel.ID,
	}

	if req.FolderID.Valid {
		if _, err := portalModel.Folders(models.FolderWhere.ID.EQ(req.FolderID.Int64)).One(ctx, boil.GetContextDB()); err != nil {
			return nil, ErrInvalidFolderID
		}
		newDoc.FolderID = req.FolderID
	} else if req.OwnerID.Valid {
		if req.OwnerID.Int64 != userModel.ID {
			return nil, ErrCannotCreateDocumentForAnotherUser
		}
		newDoc.OwnerID = req.OwnerID
	}

	err := newDoc.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, err
	}

	if err = newDoc.Reload(ctx, boil.GetContextDB()); err != nil {
		return nil, err
	}

	return newDoc, nil
}

func DocumentUpdate(ctx context.Context, documentModel *models.Document, req *UpdateDocumentRequest, userID int64) (*models.Document, error) {
	updated := *documentModel

	if req.Title != nil {
		updated.Title = *req.Title
	}

	if req.FolderID != nil {
		if !req.FolderID.Valid {
			return nil, ErrUpdateDocumentInvalidFolder
		}

		if _, err := models.Folders(models.FolderWhere.ID.EQ(req.FolderID.Int64), models.FolderWhere.PortalID.EQ(documentModel.PortalID)).One(ctx, boil.GetContextDB()); err != nil {
			return nil, ErrUpdateDocumentInvalidFolder
		}

		updated.FolderID = *req.FolderID
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

func DocumentDelete(ctx context.Context, documentModel *models.Document) error {
	_, err := documentModel.Delete(ctx, boil.GetContextDB())

	return err
}
