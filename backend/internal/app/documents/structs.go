package documents

import (
	"time"

	"github.com/aarondl/null/v8"
)

type Document struct {
	ID        int64
	FolderID  null.Int64
	OwnerID   null.Int64
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateDocumentRequest struct {
	OwnerID  null.Int64 `json:"owner_id,omitempty"`
	FolderID null.Int64 `json:"folder_id,omitempty"`
	Title    string     `json:"title"`
}

type UpdateDocumentRequest struct {
	FolderID *null.Int64 `json:"folder_id,omitempty"`
	Title    *string     `json:"title,omitempty"`
}

type DocumentResponse struct {
	ID        int64      `json:"id"`
	FolderID  null.Int64 `json:"folder_id,omitempty"`
	OwnerID   null.Int64 `json:"owner_id,omitempty"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
