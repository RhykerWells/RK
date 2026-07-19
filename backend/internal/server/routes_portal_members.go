package server

import (
	"github.com/RhykerWells/RK/backend/internal/permissions"
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/go-chi/chi/v5"
)

func registerPortalMemberRoutes(r chi.Router) {
	r.Route("/members", func(r chi.Router) {
		r.With(middleware.WithPortalMembershipMW).Get("/", handlers.PortalMembers)

		r.Route("/@me", func(r chi.Router) {
			r.Put("/", handlers.PortalMemberJoin)

			r.Group(func(r chi.Router) {
				r.Use(middleware.WithPortalMembershipMW)

				r.Get("/", handlers.PortalMember)
				r.Delete("/", handlers.PortalMemberLeave)
			})
		})

		r.Route("/{user_id}", func(r chi.Router) {
			r.Use(middleware.WithPortalMembershipMW)

			r.Get("/", handlers.PortalMember)

			r.Group(func(r chi.Router) {
				r.Use(middleware.WithPermissionsMW(permissions.PermissionPortalManageMembers))

				r.Put("/", handlers.PortalMemberAdd)
				r.Delete("/", handlers.PortalMemberRemove)
			})
			r.With(middleware.WithPermissionsMW(permissions.PermissionPortalManageMembers, permissions.PermissionPortalManageMemberRoles)).Patch("/", handlers.PortalMemberUpdate)
		})
	})
}
