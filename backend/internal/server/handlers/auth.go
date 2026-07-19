package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RhykerWells/RK/backend/internal/app/users"
	"github.com/RhykerWells/RK/backend/internal/auth"
	"github.com/RhykerWells/RK/backend/internal/config"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/RhykerWells/RK/backend/internal/server/response"
	"github.com/aarondl/null/v8"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/oauth2"
)

var OauthConf *oauth2.Config

func InitDiscordOauth() {
	OauthConf = &oauth2.Config{
		ClientID:     config.AppConfig.Discord.BotClientID,
		ClientSecret: config.AppConfig.Discord.BotSecret,
		Scopes:       []string{"identify", "guilds", "email"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://discord.com/api/oauth2/token",
			AuthURL:  "https://discord.com/api/oauth2/authorize",
		},
	}

	if config.AppConfig.Server.EnabledHTTPS {
		OauthConf.RedirectURL = "https://" + config.AppConfig.Server.Host + "/confirm"
	} else {
		OauthConf.RedirectURL = "http://" + config.AppConfig.Server.Host + "/confirm"
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	// If the user already has a valid session, send them home.
	if cookie, err := r.Cookie(auth.SessionCookieName); err == nil {
		if _, err := auth.ValidateSessionToken(r.Context(), cookie.Value); err == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	// Generate and store a CSRF token for the OAuth flow.
	csrfToken, err := createCSRF()
	if err != nil {
		http.Redirect(w, r, "?error=failed_to_create_csrf", http.StatusSeeOther)
		return
	}

	setCSRFCookie(w, csrfToken)

	// Redirect the user to Discord's OAuth2 authorization page.
	url := OauthConf.AuthCodeURL(csrfToken, oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Callback handles the successful Discord Oauth login and redirects users to the frontend
func Callback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Validate CSRF state.
	csrf := getCSRF(w, r)
	state := r.URL.Query().Get("state")
	if state != csrf {
		http.Redirect(w, r, "/?error=invalid_CSRF", http.StatusSeeOther)
		return
	}

	// Exchange the authorization code.
	code := r.FormValue("code")
	oauthToken, err := OauthConf.Exchange(ctx, code)
	if err != nil {
		http.Redirect(w, r, "/?error=oauth2_failure", http.StatusSeeOther)
		return
	}

	// Retrieving the user's Discord information using the access token.
	client := OauthConf.Client(ctx, oauthToken)
	resp, err := client.Get("https://discord.com/api/v10/users/@me")
	if err != nil {
		http.Redirect(w, r, "/?error=failed_retrieving_info", http.StatusSeeOther)
		return
	}
	defer resp.Body.Close()

	// Decoding the Discord user information from the response
	var dgoUser *discordgo.User
	if err := json.NewDecoder(resp.Body).Decode(&dgoUser); err != nil {
		http.Redirect(w, r, "/?error=failed_decoding_info", http.StatusSeeOther)
		return
	}

	// Check if the user already exists in the database by their Discord ID. If not, create a new user.
	user, err := users.GetUserByDiscordID(ctx, dgoUser.ID)
	if err != nil {
		req := &users.CreateUserRequest{
			Username:    dgoUser.Username,
			DiscordID:   null.StringFrom(dgoUser.ID),
			DisplayName: dgoUser.GlobalName,
			Email:       null.StringFrom(dgoUser.Email),
			AvatarURL:   null.StringFrom(dgoUser.AvatarURL("1024")),
		}

		user, err = users.UserCreate(ctx, req, users.AuthTypeDiscord)
		if err != nil {
			http.Redirect(w, r, "/?error=failed_creating_user", http.StatusSeeOther)
			return
		}
	}

	// Create a session for the authenticated user.
	_, sessionToken, err := auth.CreateSession(ctx, user)
	if err != nil {
		http.Redirect(w, r, "/?error=failed_creating_session", http.StatusSeeOther)
		return
	}

	auth.SetSessionCookie(w, sessionToken)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// If a session cookie exists, delete the session.
	if cookie, err := r.Cookie(auth.SessionCookieName); err == nil {
		_ = auth.DeleteSession(r.Context(), cookie.Value)
	}

	// Remove the cookie from the browser.
	auth.ClearSessionCookie(w)

	// Redirect back to the homepage (or login page).
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func IssueAPIToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.SessionFromContext(ctx)
	if !ok {
		response.ErrorMessage(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var req struct {
		ExpiresIn int64 `json:"expires_in,omitempty"` // Optional: seconds until expiration
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ErrorMessage(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Create API token
	_, token, err := auth.CreateAPIToken(ctx, session.UserID, nil)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrUserHasAPIToken):
			response.ErrorMessage(w, http.StatusBadRequest, "user already has an API token")
		default:
			response.ErrorMessage(w, http.StatusInternalServerError, "failed to create token")
		}
		return
	}

	response.JSON(w, http.StatusCreated, map[string]any{
		"token": token,
	})
}

func RevokeAPIToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.SessionFromContext(ctx)
	if !ok {
		response.ErrorMessage(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	ok, err := auth.UserHasAPIToken(ctx, session.UserID)
	if err != nil {
		response.ErrorMessage(w, http.StatusInternalServerError, "internal server error")
		return
	}

	if !ok {
		response.ErrorMessage(w, http.StatusNotFound, "user does not have an API token")
		return
	}

	// Revoke the token
	if err := auth.RevokeToken(ctx, session.UserID); err != nil {
		response.ErrorMessage(w, http.StatusInternalServerError, "internal server error")
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
