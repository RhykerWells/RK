package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RhykerWells/RK/backend/internal/app/portals"
	. "github.com/RhykerWells/RK/backend/internal/server/errors"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/RhykerWells/RK/backend/internal/server/response"
	"goji.io/v3/pat"
)

func PortalRoles(w http.ResponseWriter, r *http.Request) {
	portalModel, _ := middleware.PortalFromContext(r.Context())

	roles := portals.PortalRolesFromModel(portalModel)

	response.JSON(w, http.StatusOK, map[string]any{
		"roles": roles,
	})
}

func PortalRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	roleIDStr := pat.Param(r, "role_id")
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, ErrInvalidRoleID)
		return
	}

	role, err := portals.GetPortalRoleByID(ctx, portalModel, roleID)
	if err != nil {
		response.Error(w, http.StatusNotFound, ErrRoleNotFound)
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"role": portals.PortalRoleFromModel(role),
	})
}

func PortalRoleCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	var req portals.CreatePortalRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	role, err := portals.PortalRoleCreate(ctx, portalModel, &req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, map[string]any{
		"role": portals.PortalRoleFromModel(role),
	})
}

func PortalRoleUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	roleIDStr := pat.Param(r, "role_id")
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, ErrInvalidRoleID)
		return
	}

	role, err := portals.GetPortalRoleByID(ctx, portalModel, roleID)
	if err != nil {
		response.Error(w, http.StatusNotFound, ErrRoleNotFound)
		return
	}

	var update portals.UpdatePortalRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	role, err = portals.PortalRoleUpdate(ctx, role, &update)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"role": portals.PortalRoleFromModel(role),
	})
}

func PortalRoleDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	roleIDStr := pat.Param(r, "role_id")
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, ErrInvalidRoleID)
		return
	}

	role, err := portals.GetPortalRoleByID(ctx, portalModel, roleID)
	if err != nil {
		response.Error(w, http.StatusNotFound, ErrRoleNotFound)
		return
	}

	err = portals.PortalRoleDelete(ctx, role)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
