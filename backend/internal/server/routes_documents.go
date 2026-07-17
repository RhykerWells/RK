package server

import (
	"github.com/RhykerWells/RK/backend/internal/permissions"
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"goji.io/v3"
	"goji.io/v3/pat"
)

func registerFolderDocumentRoutes(api *goji.Mux) {
	portalRequiredMux := goji.SubMux()
	portalRequiredMux.Use(middleware.WithPortalMW)

	// Folder Endpoints
	portalRequiredMux.HandleFunc(pat.Get(EndpointPortalFolders), handlers.Folders)
	portalRequiredMux.HandleFunc(pat.Get(EndpointPortalFolder), handlers.Folder)
	portalRequiredMux.Handle(pat.Post(EndpointPortalFolders), middleware.WithPermissionsMW(handlers.FolderCreate, permissions.PermissionFolderCreate))
	portalRequiredMux.Handle(pat.Patch(EndpointPortalFolder), middleware.WithPermissionsMW(handlers.FolderUpdate, permissions.PermissionFolderEdit))
	portalRequiredMux.Handle(pat.Delete(EndpointPortalFolder), middleware.WithPermissionsMW(handlers.FolderDelete, permissions.PermissionFolderDelete))

	// Document Endpoints
	portalRequiredMux.HandleFunc(pat.Get(EndpointFolderDocument), handlers.Document)
	portalRequiredMux.Handle(pat.Post(EndpointFolderDocuments), middleware.WithPermissionsMW(handlers.DocumentCreate, permissions.PermissionDocumentCreate))
	portalRequiredMux.Handle(pat.Patch(EndpointFolderDocument), middleware.WithPermissionsMW(handlers.DocumentUpdate, permissions.PermissionDocumentEdit))
	portalRequiredMux.Handle(pat.Delete(EndpointFolderDocument), middleware.WithPermissionsMW(handlers.DocumentDelete, permissions.PermissionDocumentDelete))

	api.Handle(pat.New("/*"), portalRequiredMux)
}
