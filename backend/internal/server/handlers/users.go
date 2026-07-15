package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RhykerWells/RK/backend/internal/app/users"
	"goji.io/v3/pat"
)

var (
	ErrInvalidUserID = errors.New("invalid user id")
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidMissingRequestType = errors.New("invalid or missing request type")
)

func User(w http.ResponseWriter, r *http.Request) {
	userID := pat.Param(r, "user_id")

	ctx := r.Context()

	user, err := getUser(ctx, r.URL.Query().Get("type"), userID)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidUserID), errors.Is(err, ErrInvalidMissingRequestType):
			RespondError(w, http.StatusBadRequest, err)
		case errors.Is(err, sql.ErrNoRows):
			RespondError(w, http.StatusNotFound, ErrUserNotFound)
		default:
			RespondErrorMessage(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	returnedUser := users.UserModelToResponse(user)

	RespondJSON(w, http.StatusOK, map[string]any{
		"user": returnedUser,
	})
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var req users.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	user, err := users.UserCreate(r.Context(), &req)
	if err != nil {
		RespondErrorMessage(w, http.StatusInternalServerError, "internal server error")
		return
	}

	returnedUser := users.UserModelToResponse(user)

	RespondJSON(w, http.StatusCreated, map[string]any{
		"user":   returnedUser,
	})
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	userID := pat.Param(r, "user_id")

	ctx := r.Context()

	user, err := getUser(ctx, r.URL.Query().Get("type"), userID)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidUserID), errors.Is(err, ErrInvalidMissingRequestType):
			RespondError(w, http.StatusBadRequest, err)
		case errors.Is(err, sql.ErrNoRows):
			RespondError(w, http.StatusNotFound, ErrUserNotFound)
		default:
			RespondErrorMessage(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	err = users.UserDelete(ctx, user)
	if err != nil {
		RespondErrorMessage(w, http.StatusInternalServerError, "internal server error")
		return
	}

	RespondJSON(w, http.StatusNoContent, nil)
}
