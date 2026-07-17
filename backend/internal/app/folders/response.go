package folders

import (
	"github.com/RhykerWells/RK/backend/internal/database/models"
)

func PortalFoldersFromModel(model *models.Portal) []FolderResponse {
	folders := make([]FolderResponse, 0, len(model.GetFolders()))

	for _, folderModel := range model.GetFolders() {
		folders = append(folders, PortalFolderFromModel(folderModel))
	}

	return folders
}

func PortalFolderFromModel(model *models.Folder) FolderResponse {
	return FolderResponse{
		ID:               model.ID,
		PortalID:         model.PortalID,
		ParentFolderID:   model.ParentFolderID,
		Name:             model.Name,
		ChildFolderCount: int64(len(model.GetParentFolderFolders())),
		DocumentCount:    int64(len(model.GetDocuments())),
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
	}
}
