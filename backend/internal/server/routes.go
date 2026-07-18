package server

import (
	"net/http"

	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"github.com/RhykerWells/RK/backend/internal/server/response"
)

func (s *Server) registerRoutes() {
	handlers.InitDiscordOauth()

	s.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	s.Router.Get("/login", handlers.Login)
	s.Router.Get("/callback", handlers.Callback)
	s.Router.Get("/logout", handlers.Logout)

	s.registerUserRoutes()

	s.registerPortalRoutes()
}
