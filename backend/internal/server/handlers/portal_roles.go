package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RhykerWells/RK/backend/internal/app/portals"
	. "github.com/RhykerWells/RK/backend/internal/server/errors"
	"goji.io/v3/pat"
)

func PortalRoles(w http.ResponseWriter, r *http.Request) {
	portalModel, _ := PortalFromContext(r.Context())

	roles := portals.PortalRolesFromModel(portalModel)

	RespondJSON(w, http.StatusOK, map[string]any{
		"roles": roles,
	})
}

func PortalRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := PortalFromContext(ctx)

	roleIDStr := pat.Param(r, "role_id")
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		RespondError(w, http.StatusBadRequest, ErrInvalidRoleID)
		return
	}

	role, err := portals.GetPortalRoleByID(ctx, portalModel, roleID)
	if err != nil {
		RespondError(w, http.StatusNotFound, ErrRoleNotFound)
		return
	}

	RespondJSON(w, http.StatusOK, map[string]any{
		"role": portals.PortalRoleFromModel(role),
	})
}

func PortalRoleCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := PortalFromContext(ctx)

	var req portals.CreatePortalRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	role, err := portals.PortalRoleCreate(ctx, portalModel, &req)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	RespondJSON(w, http.StatusCreated, map[string]any{
		"role": portals.PortalRoleFromModel(role),
	})
}

func PortalRoleUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := PortalFromContext(ctx)

	roleIDStr := pat.Param(r, "role_id")
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		RespondError(w, http.StatusBadRequest, ErrInvalidRoleID)
		return
	}

	role, err := portals.GetPortalRoleByID(ctx, portalModel, roleID)
	if err != nil {
		RespondError(w, http.StatusNotFound, ErrRoleNotFound)
		return
	}

	var update portals.UpdatePortalRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	role, err = portals.PortalRoleUpdate(ctx, role, &update)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	RespondJSON(w, http.StatusOK, map[string]any{
		"role": portals.PortalRoleFromModel(role),
	})
}

func PortalRoleDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := PortalFromContext(ctx)

	roleIDStr := pat.Param(r, "role_id")
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		RespondError(w, http.StatusBadRequest, ErrInvalidRoleID)
		return
	}

	role, err := portals.GetPortalRoleByID(ctx, portalModel, roleID)
	if err != nil {
		RespondError(w, http.StatusNotFound, ErrRoleNotFound)
		return
	}

	err = portals.PortalRoleDelete(ctx, role)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
