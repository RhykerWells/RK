package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RhykerWells/RK/backend/internal/app/portals"
)

var (
	ErrInvalidPortalID = errors.New("invalid portal id")
	ErrPortalNotFound = errors.New("portal not found")
)

func Portal(w http.ResponseWriter, r *http.Request) {
	portalModel, _ := PortalFromContext(r.Context())

	returnedPortal := portals.PortalModelToResponse(portalModel)

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

	returnedPortal := portals.PortalModelToResponse(portal)

	RespondJSON(w, http.StatusCreated, map[string]any{
		"portal": returnedPortal,
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
