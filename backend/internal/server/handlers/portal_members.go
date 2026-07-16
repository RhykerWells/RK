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

func PortalMembers(w http.ResponseWriter, r *http.Request) {
	portalModel, _ := middleware.PortalFromContext(r.Context())

	members := portals.PortalMembersFromModel(portalModel)

	response.JSON(w, http.StatusOK, map[string]any{
		"members": members,
	})
}

func PortalMember(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	memberIDStr := pat.Param(r, "user_id")
	memberID, err := strconv.ParseInt(memberIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, ErrInvalidUserID)
		return
	}

	member, err := portals.GetPortalMemberByID(ctx, portalModel, memberID)
	if err != nil {
		response.Error(w, http.StatusNotFound, ErrMemberNotFound)
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"member": portals.PortalMemberFromModel(member),
	})
}

func PortalMemberCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	userIDStr := pat.Param(r, "user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, ErrInvalidUserID)
		return
	}

	member, err := portals.PortalMemberCreate(ctx, portalModel, userID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, map[string]any{
		"member": portals.PortalMemberFromModel(member),
	})
}

func PortalMemberUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	memberIDStr := pat.Param(r, "user_id")
	memberID, err := strconv.ParseInt(memberIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, ErrInvalidUserID)
		return
	}

	member, err := portals.GetPortalMemberByID(ctx, portalModel, memberID)
	if err != nil {
		response.Error(w, http.StatusNotFound, ErrMemberNotFound)
		return
	}

	var update portals.UpdatePortalMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	member, err = portals.MemberUpdate(ctx, member, &update)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"member": portals.PortalMemberFromModel(member),
	})
}

func PortalMemberDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	memberIDStr := pat.Param(r, "user_id")
	memberID, err := strconv.ParseInt(memberIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, ErrInvalidUserID)
		return
	}

	member, err := portals.GetPortalMemberByID(ctx, portalModel, memberID)
	if err != nil {
		response.Error(w, http.StatusNotFound, ErrMemberNotFound)
		return
	}

	if err := portals.MemberDelete(ctx, member); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
