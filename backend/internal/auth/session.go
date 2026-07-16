package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/RhykerWells/RK/backend/internal/config"
	"github.com/RhykerWells/RK/backend/internal/database/models"
	"github.com/aarondl/sqlboiler/v4/boil"
)

var (
	ErrInvalidSessionToken = errors.New("invalid session token")
	ErrSessionNotFound     = errors.New("session not found")
)

const sessionTokenLength = 32

const sessionDuration = 24 * 30 * time.Hour

const SessionCookieName = "rk_session"

func CreateSession(ctx context.Context, user *models.User) (*models.Session, string, error) {
	token, err := GenerateToken()
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate session token: %w", err)
	}

	session := &models.Session{
		UserID:    user.ID,
		TokenHash: HashToken(token),
		ExpiresAt: time.Now().Add(sessionDuration),
		CreatedAt: time.Now(),
	}

	if err := session.Insert(ctx, boil.GetContextDB(), boil.Infer()); err != nil {
		return nil, "", fmt.Errorf("failed to insert session: %w", err)
	}

	return session, token, nil
}

func ValidateSessionToken(ctx context.Context, token string) (*models.Session, error) {
	if token == "" {
		return nil, ErrInvalidSessionToken
	}

	tokenHash := HashToken(token)

	session, err := models.Sessions(models.SessionWhere.TokenHash.EQ(tokenHash)).One(ctx, boil.GetContextDB())
	if err != nil {
		return nil, err
	}

	// Check if session has expired
	if time.Now().After(session.ExpiresAt) {
		return nil, ErrSessionNotFound
	}

	return session, nil
}

func DeleteSession(ctx context.Context, token string) error {
	if token == "" {
		return ErrInvalidSessionToken
	}

	tokenHash := HashToken(token)

	_, err := models.Sessions(models.SessionWhere.TokenHash.EQ(tokenHash)).DeleteAll(ctx, boil.GetContextDB())
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserSessions(ctx context.Context, userID int64) error {
	_, err := models.Sessions(models.SessionWhere.UserID.EQ(userID)).DeleteAll(ctx, boil.GetContextDB())

	return err
}

func SetSessionCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:     SessionCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   config.AppConfig.Server.EnabledHTTPS,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(sessionDuration),
		MaxAge:   int(sessionDuration.Seconds()),
	}
	http.SetCookie(w, cookie)
}

func ClearSessionCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     SessionCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   config.AppConfig.Server.EnabledHTTPS,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(-time.Hour),
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
}
