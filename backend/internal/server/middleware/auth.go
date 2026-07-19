package middleware

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/RhykerWells/RK/backend/internal/app/portals"
	"github.com/RhykerWells/RK/backend/internal/app/users"
	"github.com/RhykerWells/RK/backend/internal/auth"
	"github.com/RhykerWells/RK/backend/internal/database/models"
	"github.com/RhykerWells/RK/backend/internal/permissions"
	"github.com/RhykerWells/RK/backend/internal/server/response"
)

// WithAuthMW validates requests using either a session cookie, or an API token.
// If a valid session/token is found, the user and session (where applicable) is stored on the request context.
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
				userModel, err := users.GetUserByID(ctx, session.UserID)
				if err == nil {
					ctx = context.WithValue(ctx, ContextSessionKey, session)
					ctx = context.WithValue(ctx, ContextUserKey, userModel)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				response.ErrorMessage(w, http.StatusInternalServerError, "internal server error")
				return
			}
		}

		// Try Bearer token
		// This is an external service interacting with our API
		if bearerToken != "" {
			apiToken, err := auth.ValidateAPIToken(ctx, bearerToken)
			if err == nil {
				userModel, err := users.GetUserByID(ctx, apiToken.UserID)
				if err == nil {
					ctx = context.WithValue(ctx, ContextUserKey, userModel)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				response.ErrorMessage(w, http.StatusInternalServerError, "internal server error")
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

	sessionModel, ok := v.(*models.Session)
	return sessionModel, ok
}

func UserFromContext(ctx context.Context) (*models.User, bool) {
	v := ctx.Value(ContextUserKey)
	if v == nil {
		return nil, false
	}

	userModel, ok := v.(*models.User)
	return userModel, ok
}

// RequireAdminMW checks if the current user is a site administrator and prevents access to the next handler if they are not.
func RequireAdminMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userModel, ok := UserFromContext(ctx)
		if !ok {
			response.ErrorMessage(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		if !userModel.IsAdministrator {
			response.ErrorMessage(w, http.StatusForbidden, "forbidden")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// WithPortalMembershipMW checks if the current user is a member of the portal and prevents access to the next handler if they are not.
// If the user is an administrator, they are allowed access regardless of membership.
// If a valid membership (excluding administration) is found, it is stored on the request context.
func WithPortalMembershipMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userModel, ok := UserFromContext(ctx)
		if !ok {
			response.ErrorMessage(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		if userModel.IsAdministrator {
			next.ServeHTTP(w, r)
			return
		}

		portalModel, ok := PortalFromContext(ctx)
		if !ok {
			response.ErrorMessage(w, http.StatusInternalServerError, "internal server error")
			return
		}

		memberModel, err := portals.GetPortalMemberByID(ctx, portalModel, userModel.ID)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				response.ErrorMessage(w, http.StatusForbidden, "forbidden")
			default:
				response.ErrorMessage(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		ctx = context.WithValue(ctx, ContextMemberKey, memberModel)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PortalMemberFromContext(ctx context.Context) (*models.PortalMembership, bool) {
	v := ctx.Value(ContextMemberKey)
	if v == nil {
		return nil, false
	}

	memberModel, ok := v.(*models.PortalMembership)
	return memberModel, ok
}

func WithPermissionsMW(perms ...permissions.Permission) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			userModel, ok := UserFromContext(ctx)
			if !ok {
				response.ErrorMessage(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			if userModel.IsAdministrator {
				next.ServeHTTP(w, r)
				return
			}

			memberModel, ok := PortalMemberFromContext(ctx)
			if !ok {
				response.ErrorMessage(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			if !hasRequiredPermission(portals.PortalMemberFromModel(memberModel).Roles, perms) {
				response.ErrorMessage(w, http.StatusForbidden, "forbidden")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
