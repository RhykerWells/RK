package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RhykerWells/RK/backend/internal/app/portals"
)

var (
	ErrInvalidPortalID = errors.New("invalid portal id")
)

func Portal(w http.ResponseWriter, r *http.Request) {
	portalModel, _ := PortalFromContext(r.Context())

	roles := getPortalRoles(portalModel)
	returnedPortal := &portals.PortalResponse{
		ID: portalModel.ID,
		Name: portalModel.Name,
		Domain: portalModel.Domain,
		Roles: roles,
		CreatedAt: portalModel.CreatedAt,
		UpdatedAt: portalModel.UpdatedAt,
	}

	RespondJSON(w, http.StatusOK, map[string]any{
		"portal": returnedPortal,
	})
}

func PortalCreate(w http.ResponseWriter, r *http.Request) {
	var req portals.CreatePortalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	portal, err := portals.PortalCreate(r.Context(), &req)
	if err != nil {
		RespondErrorMessage(w, http.StatusInternalServerError, "internal server error")
		return
	}

	returnedPortal := &portals.PortalResponse{
		ID: portal.ID,
		Name: portal.Name,
		Domain: portal.Domain,
		CreatedAt: portal.CreatedAt,
		UpdatedAt: portal.UpdatedAt,
	}

	RespondJSON(w, http.StatusCreated, map[string]any{
		"user":   returnedPortal,
	})
}

func PortalDelete(w http.ResponseWriter, r *http.Request) {
	portalModel, _ := PortalFromContext(r.Context())

	err := portals.PortalDelete(r.Context(), portalModel)
	if err != nil {
		RespondErrorMessage(w, http.StatusInternalServerError, "internal server error")
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}