package server

import (
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"goji.io/v3"
	"goji.io/v3/pat"
)

func (s *Server) registerRoutes() {
	handlers.InitDiscordOauth()

	api := goji.SubMux()

	api.HandleFunc(pat.Get(EndpointHealth), handlers.Health)

	registerAuthRoutes(api)

	authRequiredMux := goji.SubMux()
	authRequiredMux.Use(middleware.WithAuthMW)

	registerUserRoutes(authRequiredMux)

	registerPortalRoutes(authRequiredMux)

	api.Handle(pat.New("/*"), authRequiredMux)

	s.Multiplexer.Handle(pat.New(EndpointAPI+"/*"), api)
}
