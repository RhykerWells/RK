package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RhykerWells/RK/backend/internal/app/portals"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/RhykerWells/RK/backend/internal/server/response"
)

func Portals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	portalModels, err := portals.GetPortals(ctx)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"portals": portals.PortalsModelToResponse(portalModels),
	})
}

func Portal(w http.ResponseWriter, r *http.Request) {
	portalModel, _ := middleware.PortalFromContext(r.Context())

	response.JSON(w, http.StatusOK, map[string]any{
		"portal": portals.PortalModelToResponse(portalModel),
	})
}

func PortalCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

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

	response.JSON(w, http.StatusCreated, map[string]any{
		"portal": portals.PortalModelToResponse(portal),
	})
}

func PortalUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	portalModel, _ := middleware.PortalFromContext(ctx)

	var req portals.UpdatePortalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	portal, err := portals.PortalUpdate(ctx, portalModel, &req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, map[string]any{
		"portal": portals.PortalModelToResponse(portal),
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
