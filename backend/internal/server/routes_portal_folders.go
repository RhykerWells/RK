package server

import (
	"github.com/RhykerWells/RK/backend/internal/permissions"
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/go-chi/chi/v5"
)

func registerPortalFolderRoutes(r chi.Router) {
	r.Route("/folders", func(r chi.Router) {
		r.Use(middleware.WithPortalMembershipMW)

		r.Get("/", handlers.PortalFolders)

		r.Group(func(r chi.Router) {
			r.Use(middleware.WithPermissionsMW(permissions.PermissionPortalManageFolders))

			r.Post("/", handlers.PortalFolderCreate)
			r.Patch("/{folder_id}", handlers.PortalFolderUpdate)
			r.Delete("/{folder_id}", handlers.PortalFolderDelete)
		})

		r.Get("/{folder_id}", handlers.PortalFolder)
	})
}
