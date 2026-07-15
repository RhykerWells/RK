package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strconv"
	"time"

	"github.com/RhykerWells/RK/backend/internal/app/users"
	"github.com/RhykerWells/RK/backend/internal/database/models"
	. "github.com/RhykerWells/RK/backend/internal/server/errors"
)

const (
	CSRFToken = "rk_csrf"
)

func createCSRF() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func setCSRFCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     CSRFToken,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(300 * time.Second),
	})
}

func getCSRF(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie(CSRFToken)
	if err == nil {
		return cookie.Value
	}

	http.SetCookie(w, &http.Cookie{
		Name:     CSRFToken,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Unix(0, 0),
	})
	return ""
}

func getUser(ctx context.Context, requestType string, id string) (*models.User, error) {
	var (
		user *models.User
		err  error
	)

	switch requestType {
	case "id":
		id, convErr := strconv.ParseInt(id, 10, 64)
		if convErr != nil {
			return nil, ErrInvalidUserID
		}
		user, err = users.GetUserByID(ctx, id)
	case "discord":
		user, err = users.GetUserByDiscordID(ctx, id)
	default:
		return nil, ErrInvalidMissingRequestType
	}

	return user, err
}
