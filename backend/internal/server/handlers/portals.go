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
