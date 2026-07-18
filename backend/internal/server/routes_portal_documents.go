package server

import (
	"github.com/RhykerWells/RK/backend/internal/permissions"
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/go-chi/chi/v5"
)

func registerPortalDocumentRoutes(r chi.Router) {
	r.Route("/documents", func(r chi.Router) {
		r.Use(middleware.WithPortalMembershipMW)

		r.Get("/", handlers.PortalDocuments)

		r.Group(func(r chi.Router) {
			r.Use(middleware.WithPermissionsMW(permissions.PermissionPortalManageDocuments))

			r.Post("/", handlers.PortalDocumentCreate)
			r.Patch("/{document_id}", handlers.PortalDocumentUpdate)
			r.Delete("/{document_id}", handlers.PortalDocumentDelete)
		})

		r.Get("/{document_id}", handlers.PortalDocument)
	})
}
