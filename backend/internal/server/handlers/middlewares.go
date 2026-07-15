package handlers

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/RhykerWells/RK/backend/internal/app/portals"
	"github.com/RhykerWells/RK/backend/internal/database/models"
	"goji.io/v3/pat"
)

// PortalContextKey is the context key used to store the loaded portal.
type PortalContextKey string

const ContextPortalKey PortalContextKey = "portal"

// WithPortalMW loads a portal by the path parameter ("portal_id" or "id")
// and stores it in the request context under ContextPortalKey. If the portal cannot
// be loaded the middleware writes an error response and aborts the chain.
func WithPortalMW(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        portalIDStr := pat.Param(r, "portal_id")

        ctx := r.Context()

        portalIDInt, err := strconv.ParseInt(portalIDStr, 10, 64)
        if err != nil {
			RespondErrorMessage(w, http.StatusBadRequest, err.Error())
            return
        }

        portalModel, err := portals.GetPortalByID(ctx, portalIDInt)
        if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				RespondErrorMessage(w, http.StatusNotFound, "portal not found")
			default:
				RespondErrorMessage(w, http.StatusInternalServerError, "internal server error")
			}
			return
        }

        ctx = context.WithValue(ctx, ContextPortalKey, portalModel)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// PortalFromContext is a helper to retrieve the loaded portal from context.
func PortalFromContext(ctx context.Context) (*models.Portal, bool) {
    v := ctx.Value(ContextPortalKey)
    if v == nil {
        return nil, false
    }
    p, ok := v.(*models.Portal)
    return p, ok
}
