package server

import (
	"github.com/RhykerWells/RK/backend/internal/permissions"
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/go-chi/chi/v5"
)

func registerPortalRoleRoutes(r chi.Router) {
	r.Route("/roles", func(r chi.Router) {
		r.Use(middleware.WithPortalMembershipMW)

		r.Get("/", handlers.PortalRoles)

		r.With(middleware.WithPermissionsMW(permissions.PermissionPortalManageRoles)).Post("/", handlers.PortalRoleCreate)

		r.Route("/{role_id}", func(r chi.Router) {
			r.Get("/", handlers.PortalRole)

			r.Group(func(r chi.Router) {
				r.Use(middleware.WithPermissionsMW(permissions.PermissionPortalManageRoles))

				r.Patch("/", handlers.PortalRoleUpdate)
				r.Delete("/", handlers.PortalRoleDelete)
			})
		})
	})
}
