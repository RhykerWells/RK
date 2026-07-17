package folders

import (
	"time"

	"github.com/RhykerWells/RK/backend/internal/app/documents"
	"github.com/aarondl/null/v8"
)

type Folder struct {
	ID             int64
	PortalID       int64
	ParentFolderID null.Int64
	Name           string
	CreatedBy      int64
	UpdatedBy      null.Int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type FolderResponse struct {
	ID               int64      `json:"id"`
	PortalID         int64      `json:"portal_id"`
	ParentFolderID   null.Int64 `json:"parent_folder_id,omitempty"`
	Name             string     `json:"name"`
	ChildFolderCount int64      `json:"child_folder_count"`
	DocumentCount    int64      `json:"document_count"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

type FolderContentsResponse struct {
	Folders   []FolderResponse             `json:"folders,omitempty"`
	Documents []documents.DocumentResponse `json:"documents,omitempty"`
}

type CreateFolderRequest struct {
	Name           string `json:"name"`
	ParentFolderID *int64 `json:"parent_folder_id,omitempty"`
}

type UpdateFolderRequest struct {
	Name           *string     `json:"name,omitempty"`
	ParentFolderID *null.Int64 `json:"parent_folder_id,omitempty"`
}
