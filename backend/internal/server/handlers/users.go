package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RhykerWells/RK/backend/internal/app/users"
	. "github.com/RhykerWells/RK/backend/internal/server/errors"
	"github.com/RhykerWells/RK/backend/internal/server/response"
	"goji.io/v3/pat"
)

func User(w http.ResponseWriter, r *http.Request) {
	userID := pat.Param(r, "user_id")

	ctx := r.Context()

	user, err := getUser(ctx, r.URL.Query().Get("type"), userID)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidUserID), errors.Is(err, ErrInvalidMissingRequestType):
			response.Error(w, http.StatusBadRequest, err)
		case errors.Is(err, sql.ErrNoRows):
			response.Error(w, http.StatusNotFound, ErrUserNotFound)
		default:
			response.Error(w, http.StatusInternalServerError, err)
		}
		return
	}

	returnedUser := users.UserModelToResponse(user)

	response.JSON(w, http.StatusOK, map[string]any{
		"user": returnedUser,
	})
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var req users.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	user, err := users.UserCreate(r.Context(), &req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	returnedUser := users.UserModelToResponse(user)

	response.JSON(w, http.StatusCreated, map[string]any{
		"user": returnedUser,
	})
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	userID := pat.Param(r, "user_id")

	ctx := r.Context()

	user, err := getUser(ctx, r.URL.Query().Get("type"), userID)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidUserID), errors.Is(err, ErrInvalidMissingRequestType):
			response.Error(w, http.StatusBadRequest, err)
		case errors.Is(err, sql.ErrNoRows):
			response.Error(w, http.StatusNotFound, ErrUserNotFound)
		default:
			response.Error(w, http.StatusInternalServerError, err)
		}
		return
	}

	err = users.UserDelete(ctx, user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
