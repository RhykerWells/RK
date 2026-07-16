package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/RhykerWells/RK/backend/internal/app/users"
	"github.com/RhykerWells/RK/backend/internal/auth"
	"github.com/RhykerWells/RK/backend/internal/database/models"
	"github.com/RhykerWells/RK/backend/internal/server/response"
)

// WithAuthMW validates requests using either a session cookie, or an API token.
// If a valid session/token is found, the session is stored on the request context.
func WithAuthMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		cookieToken := ""
		if cookie, err := r.Cookie(auth.SessionCookieName); err == nil {
			cookieToken = strings.TrimSpace(cookie.Value)
		}

		authorizationHeader := r.Header.Get("Authorization")
		bearerToken := ""
		if strings.HasPrefix(strings.ToLower(authorizationHeader), "bearer ") {
			bearerToken = strings.TrimSpace(authorizationHeader[7:])
		}

		var session *models.Session
		var err error

		// Try session cookie first
		// This is our frontend interacting
		if cookieToken != "" {
			session, err = auth.ValidateSessionToken(ctx, cookieToken)
			if err == nil {
				user, _ := users.GetUserByID(ctx, session.UserID)
				ctx = context.WithValue(ctx, ContextSessionKey, session)
				ctx = context.WithValue(ctx, ContextUserKey, user)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		// Try Bearer token
		// This is an external service interacting with our API
		if bearerToken != "" {
			apiToken, err := auth.ValidateAPIToken(ctx, bearerToken)
			if err == nil {
				user, _ := users.GetUserByID(ctx, apiToken.UserID)
				ctx = context.WithValue(ctx, ContextUserKey, user)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		response.ErrorMessage(w, http.StatusUnauthorized, "unauthorized")
	})
}

func SessionFromContext(ctx context.Context) (*models.Session, bool) {
	v := ctx.Value(ContextSessionKey)
	if v == nil {
		return nil, false
	}

	session, ok := v.(*models.Session)
	return session, ok
}

func UserFromContext(ctx context.Context) (*models.User, bool) {
	v := ctx.Value(ContextUserKey)
	if v == nil {
		return nil, false
	}

	user, ok := v.(*models.User)
	return user, ok
}
