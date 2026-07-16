package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RhykerWells/RK/backend/internal/app/portals"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/RhykerWells/RK/backend/internal/server/response"
)

func Portal(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)
	user, _ := middleware.UserFromContext(ctx)

	// Check if user is admin or belongs to this portal
	if !user.IsAdministrator {
		_, err := portals.GetPortalMemberByID(ctx, portalModel, user.ID)
		if err != nil {
			response.ErrorMessage(w, http.StatusForbidden, "forbidden")
			return
		}
	}

	returnedPortal := portals.PortalModelToResponse(portalModel)

	response.JSON(w, http.StatusOK, map[string]any{
		"portal": returnedPortal,
	})
}

func PortalCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := middleware.UserFromContext(ctx)
	if !ok {
		response.ErrorMessage(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if !user.IsAdministrator {
		response.ErrorMessage(w, http.StatusForbidden, "forbidden")
		return
	}

	var req portals.CreatePortalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	portal, err := portals.PortalCreate(ctx, &req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	returnedPortal := portals.PortalModelToResponse(portal)

	response.JSON(w, http.StatusCreated, map[string]any{
		"portal": returnedPortal,
	})
}

func PortalDelete(w http.ResponseWriter, r *http.Request) {
	portalModel, _ := middleware.PortalFromContext(r.Context())

	err := portals.PortalDelete(r.Context(), portalModel)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
