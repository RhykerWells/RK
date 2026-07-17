package documents

import (
	"github.com/RhykerWells/RK/backend/internal/database/models"
)

func PortalDocumentsFromModel(model *models.Portal) []DocumentResponse {
	documents := make([]DocumentResponse, 0, len(model.GetDocuments()))

	for _, documentModel := range model.GetDocuments() {
		documents = append(documents, PortalDocumentFromModel(documentModel))
	}

	return documents
}

func PortalDocumentFromModel(model *models.Document) DocumentResponse {
	return DocumentResponse{
		ID:        model.ID,
		FolderID:  model.FolderID,
		OwnerID:   model.OwnerID,
		Title:     model.Title,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
