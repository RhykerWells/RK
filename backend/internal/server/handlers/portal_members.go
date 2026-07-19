package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RhykerWells/RK/backend/internal/app/portals"
	"github.com/RhykerWells/RK/backend/internal/database/models"
	. "github.com/RhykerWells/RK/backend/internal/server/errors"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/RhykerWells/RK/backend/internal/server/response"
	"goji.io/v3/pat"
)

func getPortalMemberFromPath(ctx context.Context, r *http.Request, portalModel *models.Portal, allowMe bool) (*models.PortalMembership, error) {
	userParam := pat.Param(r, "user_id")

	var userID int64

	if userParam == "@me" {
		if !allowMe {
			return nil, ErrInvalidUserID
		}
		userModel, _ := middleware.UserFromContext(ctx)
		userID = userModel.ID
	} else {
		var err error
		userID, err = strconv.ParseInt(userParam, 10, 64)
		if err != nil {
			return nil, ErrInvalidUserID
		}
	}

	member, err := portals.GetPortalMemberByID(ctx, portalModel, userID)
	if err != nil {
		return nil, ErrMemberNotFound
	}

	return member, nil
}

func PortalMembers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	members := portals.PortalMembersFromModel(portalModel)

	response.JSON(w, http.StatusOK, map[string]any{
		"members": members,
	})
}

func PortalMember(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	member, err := getPortalMemberFromPath(ctx, r, portalModel, true)
	if err != nil {
		switch err {
		case ErrInvalidUserID:
			response.Error(w, http.StatusBadRequest, err)
		case ErrMemberNotFound:
			response.Error(w, http.StatusNotFound, err)
		default:
			response.ErrorMessage(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"member": portals.PortalMemberFromModel(member),
	})
}

func PortalMemberAdd(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	userIDStr := pat.Param(r, "user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, ErrInvalidUserID)
		return
	}

	if _, err := portals.GetPortalMemberByID(ctx, portalModel, userID); err == nil {
		response.ErrorMessage(w, http.StatusBadRequest, "member already exists")
		return
	}

	member, err := portals.PortalMemberCreate(ctx, portalModel, userID)
	if err != nil {
		switch err {
		case ErrUserNotFound:
			response.Error(w, http.StatusNotFound, err)
		case ErrMemberAlreadyExists:
			response.Error(w, http.StatusBadRequest, err)
		default:
			response.Error(w, http.StatusBadRequest, err)
		}
		return
	}

	response.JSON(w, http.StatusCreated, map[string]any{
		"member": portals.PortalMemberFromModel(member),
	})
}

func PortalMemberUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	member, err := getPortalMemberFromPath(ctx, r, portalModel, true)
	if err != nil {
		switch err {
		case ErrInvalidUserID:
			response.Error(w, http.StatusBadRequest, err)
		case ErrMemberNotFound:
			response.Error(w, http.StatusNotFound, err)
		default:
			response.ErrorMessage(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	var update portals.UpdatePortalMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	member, err = portals.PortalMemberUpdate(ctx, member, &update)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"member": portals.PortalMemberFromModel(member),
	})
}

func PortalMemberRemove(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	member, err := getPortalMemberFromPath(ctx, r, portalModel, false)
	if err != nil {
		switch err {
		case ErrInvalidUserID:
			response.Error(w, http.StatusBadRequest, err)
		case ErrMemberNotFound:
			response.Error(w, http.StatusNotFound, err)
		default:
			response.ErrorMessage(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	if err := portals.MemberDelete(ctx, member); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func PortalMemberJoin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)
	userModel, _ := middleware.UserFromContext(ctx)

	member, err := portals.PortalMemberCreate(ctx, portalModel, userModel.ID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusCreated, map[string]any{
		"member": portals.PortalMemberFromModel(member),
	})
}

func PortalMemberLeave(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)
	userModel, _ := middleware.UserFromContext(ctx)

	member, err := portals.GetPortalMemberByID(ctx, portalModel, userModel.ID)
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
