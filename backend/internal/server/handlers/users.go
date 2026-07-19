package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/RhykerWells/RK/backend/internal/app/users"
	. "github.com/RhykerWells/RK/backend/internal/server/errors"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/RhykerWells/RK/backend/internal/server/response"
	"github.com/go-chi/chi/v5"
	"goji.io/v3/pat"
)

func Users(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	usersModel, err := users.GetUsers(ctx)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"users": users.UsersModelToResponse(usersModel),
	})
}

func User(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := chi.URLParam(r, "user_id")
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

	response.JSON(w, http.StatusOK, map[string]any{
		"user": users.UserModelToResponse(user),
	})
}

func Me(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, _ := middleware.UserFromContext(ctx)

	response.JSON(w, http.StatusOK, map[string]any{
		"user": users.UserModelToResponse(user),
	})
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := pat.Param(r, "user_id")
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
