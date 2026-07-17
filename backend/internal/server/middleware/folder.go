package middleware

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/RhykerWells/RK/backend/internal/app/folders"
	"github.com/RhykerWells/RK/backend/internal/database/models"
	. "github.com/RhykerWells/RK/backend/internal/server/errors"
	"github.com/RhykerWells/RK/backend/internal/server/response"
	"goji.io/v3/pat"
)

// WithFolderMW loads a folder by the path parameter ("folder_id")
// and stores it in the request context under ContextFolderKey. If the folder cannot
// be loaded the middleware writes an error response and aborts the chain.
func WithFolderMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		portalModel, _ := PortalFromContext(ctx)
		folderIDStr := pat.Param(r, "folder_id")

		folderIDInt, err := strconv.ParseInt(folderIDStr, 10, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, ErrInvalidFolderID)
			return
		}

		folderModel, err := folders.GetPortalFolderByID(ctx, portalModel, folderIDInt)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				response.Error(w, http.StatusNotFound, ErrFolderNotFound)
			default:
				response.ErrorMessage(w, http.StatusInternalServerError, "internal server error")
			}
		}

		ctx = context.WithValue(ctx, ContextFolderKey, folderModel)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func FolderFromContext(ctx context.Context) (*models.Folder, bool) {
	v := ctx.Value(ContextFolderKey)
	if v == nil {
		return nil, false
	}

	f, ok := v.(*models.Folder)
	return f, ok
}
