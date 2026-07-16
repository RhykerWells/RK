package server

import (
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"goji.io/v3"
	"goji.io/v3/pat"
)

func registerAuthRoutes(api *goji.Mux) {
	// Public
	api.HandleFunc(pat.Get(EndpointAuthLogin), handlers.Login)
	api.HandleFunc(pat.Get(EndpointAuthCallback), handlers.Callback)
	api.HandleFunc(pat.Get(EndpointAuthLogout), handlers.Logout)
}
