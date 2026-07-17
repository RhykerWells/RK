package middleware

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/RhykerWells/RK/backend/internal/app/documents"
	. "github.com/RhykerWells/RK/backend/internal/server/errors"
	"github.com/RhykerWells/RK/backend/internal/server/response"
	"goji.io/v3/pat"
)

// WithDocumentMw loads a document by the path parameter ("document_id")
// and stores it in the request context under ContextDocumentKey. If the document cannot
// be loaded the middleware writes an error response and aborts the chain.
func WithDocumentMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		documentIDStr := pat.Param(r, "document_id")

		documentIDInt, err := strconv.ParseInt(documentIDStr, 10, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, ErrInvalidDocumentID)
			return
		}

		documentModel, err := documents.GetDocumentByID(ctx, documentIDInt)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				response.Error(w, http.StatusNotFound, ErrDocumentNotFound)
			default:
				response.ErrorMessage(w, http.StatusInternalServerError, "internal server error")
			}
		}

		ctx = context.WithValue(ctx, ContextDocumentKey, documentModel)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
