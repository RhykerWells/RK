package documents

import (
	"context"
	"time"

	"github.com/RhykerWells/RK/backend/internal/database/models"
	. "github.com/RhykerWells/RK/backend/internal/server/errors"
	"github.com/aarondl/sqlboiler/v4/boil"
)

func GetDocumentByID(ctx context.Context, id int64) (*models.Document, error) {
	d, e := models.FindDocument(ctx, boil.GetContextDB(), id)

	return d, e
}

func DocumentCreate(ctx context.Context, req *CreateDocumentRequest, userID int64) (*models.Document, error) {
	newDoc := &models.Document{
		Title:     req.Title,
		CreatedBy: userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if req.FolderID.Valid == req.OwnerID.Valid {
		return nil, ErrNewDocumentInvalidLocation
	}

	switch {
	case req.FolderID.Valid:
		newDoc.FolderID = req.FolderID
	case req.OwnerID.Valid:
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

func DocumentUpdate(ctx context.Context, document *models.Document, req *UpdateDocumentRequest, userID int64) (*models.Document, error) {
	updated := *document

	if req.Title != nil {
		updated.Title = *req.Title
	}

	if req.FolderID != nil {
		if !req.FolderID.Valid {
			return nil, ErrUpdateDocumentInvalidFolder
		}

		updated.FolderID = *req.FolderID
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

func DocumentDelete(ctx context.Context, document *models.Document) error {
	_, err := document.Delete(ctx, boil.GetContextDB())

	return err
}
