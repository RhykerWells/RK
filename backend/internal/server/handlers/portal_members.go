package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/RhykerWells/RK/backend/internal/app/portals"
	"goji.io/v3/pat"
)

var (
	ErrMemberNotFound = errors.New("member not found")
)

func PortalMembers(w http.ResponseWriter, r *http.Request) {
	portalModel, _ := PortalFromContext(r.Context())

	members := portals.PortalMembersFromModel(portalModel)

	RespondJSON(w, http.StatusOK, map[string]any{
		"members": members,
	})
}

func PortalMember(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := PortalFromContext(ctx)

	memberIDStr := pat.Param(r, "user_id")
	memberID, err := strconv.ParseInt(memberIDStr, 10, 64)
	if err != nil {
		RespondError(w, http.StatusBadRequest, ErrInvalidUserID)
		return
	}

	member, err := portals.GetPortalMemberByID(ctx, portalModel, memberID)
	if err != nil {
		RespondError(w, http.StatusNotFound, ErrMemberNotFound)
		return
	}

	RespondJSON(w, http.StatusOK, map[string]any{
		"member": portals.PortalMemberFromModel(member),
	})
}

func PortalMemberCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := PortalFromContext(ctx)

	memberIDStr := pat.Param(r, "user_id")
	memberID, err := strconv.ParseInt(memberIDStr, 10, 64)
	if err != nil {
		RespondError(w, http.StatusBadRequest, ErrInvalidUserID)
		return
	}

	member, err := portals.PortalMemberCreate(ctx, portalModel, memberID)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	RespondJSON(w, http.StatusCreated, map[string]any{
		"member": portals.PortalMemberFromModel(member),
	})
}

func PortalMemberUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := PortalFromContext(ctx)

	memberIDStr := pat.Param(r, "user_id")
	memberID, err := strconv.ParseInt(memberIDStr, 10, 64)
	if err != nil {
		RespondError(w, http.StatusBadRequest, ErrInvalidUserID)
		return
	}

	// Todo user verification (check if found)

	var update portals.UpdatePortalMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	member, err := portals.MemberUpdate(ctx, portalModel, memberID, &update)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	RespondJSON(w, http.StatusOK, map[string]any{
		"member": portals.PortalMemberFromModel(member),
	})
}

func PortalMemberDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := PortalFromContext(ctx)

	memberIDStr := pat.Param(r, "user_id")
	memberID, err := strconv.ParseInt(memberIDStr, 10, 64)
	if err != nil {
		RespondError(w, http.StatusBadRequest, ErrInvalidUserID)
		return
	}

	if err := portals.MemberDelete(ctx, portalModel, memberID); err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
