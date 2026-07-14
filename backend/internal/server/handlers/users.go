package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RhykerWells/RK/backend/internal/app/users"
	"github.com/RhykerWells/RK/backend/internal/database/models"
	"goji.io/v3/pat"
)

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := pat.Param(r, "id")

	var (
		user *models.User
		err  error
	)

	ctx := r.Context()

	switch {
	case r.URL.Query().Get("type") == "id":
		id, convErr := strconv.ParseInt(userID, 10, 64)
		if convErr != nil {
			http.Error(w, "invalid user id", http.StatusBadRequest)
			return
		}
		user, err = users.GetUserByID(ctx, id)
	case r.URL.Query().Get("type") == "discord":
		user, err = users.GetUserByDiscordID(ctx, userID)
	default:
		http.Error(w, "missing search identifier", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	returnedUser := &users.UserResponse{
		ID:              user.ID,
		AuthType:        user.AuthType,
		DiscordID:       user.DiscordID,
		Username:        user.Username,
		DisplayName:     user.DisplayName,
		Email:           user.Email,
		AvatarURL:       user.AvatarURL,
		IsAdministrator: user.IsAdministrator,
	}

	if user.LastLoginAt.Valid {
		t := user.LastLoginAt.Time
		returnedUser.LastLoginAt = &t
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
		"user":   returnedUser,
	})
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req users.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := users.UserCreate(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	returnedUser := &users.UserResponse{
		ID:              user.ID,
		AuthType:        user.AuthType,
		Username:        user.Username,
		DiscordID:       user.DiscordID,
		DisplayName:     user.DisplayName,
		Email:           user.Email,
		AvatarURL:       user.AvatarURL,
		IsAdministrator: user.IsAdministrator,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
		"user":   returnedUser,
	})
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := pat.Param(r, "id")

	var (
		user *models.User
		err  error
	)

	ctx := r.Context()

	switch {
	case r.URL.Query().Get("type") == "id":
		id, convErr := strconv.ParseInt(userID, 10, 64)
		if convErr != nil {
			http.Error(w, "invalid user id", http.StatusBadRequest)
			return
		}
		user, err = users.GetUserByID(ctx, id)
	case r.URL.Query().Get("type") == "discord":
		user, err = users.GetUserByDiscordID(ctx, userID)
	default:
		http.Error(w, "missing search identifier", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = users.UserDelete(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
