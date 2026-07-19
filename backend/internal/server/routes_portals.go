package server

import (
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/go-chi/chi/v5"
)

func (s *Server) registerPortalRoutes() {
	s.Router.Route("/portals", func(r chi.Router) {
		r.Use(middleware.WithAuthMW)

		r.Get("/", handlers.Portals)

		r.With(middleware.RequireAdminMW).Post("/", handlers.PortalCreate)

		r.Route("/{portal_id}", func(r chi.Router) {
			r.Use(middleware.WithPortalMW)

			r.Get("/", handlers.Portal)

			r.Group(func(r chi.Router) {
				r.Use(middleware.RequireAdminMW)

				r.Put("/", handlers.PortalUpdate)

				r.Delete("/", handlers.PortalDelete)
			})

			registerPortalRoleRoutes(r)
			registerPortalFolderRoutes(r)
			registerPortalDocumentRoutes(r)
			registerPortalMemberRoutes(r)
		})
	})
}
