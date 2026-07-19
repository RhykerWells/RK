package server

import (
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/go-chi/chi/v5"
)

func (s *Server) registerUserRoutes() {
	s.Router.Route("/users", func(r chi.Router) {
		r.Use(middleware.WithAuthMW)

		r.Get("/", handlers.Users)

		r.Route("/{user_id}", func(r chi.Router) {
			r.Get("/", handlers.User)

			r.With(middleware.RequireAdminMW).Delete("/", handlers.UserDelete)
		})
		r.Route("/@me", func(r chi.Router) {
			r.Get("/", handlers.Me)

			r.Route("/api", func(r chi.Router) {
				r.Post("/", handlers.IssueAPIToken)
				r.Delete("/", handlers.RevokeAPIToken)
			})
		})
	})
}
