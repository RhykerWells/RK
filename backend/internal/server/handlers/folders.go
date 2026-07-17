package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RhykerWells/RK/backend/internal/app/folders"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/RhykerWells/RK/backend/internal/server/response"
	"goji.io/v3/pat"
)

func Folders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	folders := folders.PortalFoldersFromModel(portalModel)

	response.JSON(w, http.StatusOK, map[string]any{
		"folders": folders,
	})
}

func Folder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	folderIDStr := pat.Param(r, "folder_id")
	folderID, err := strconv.ParseInt(folderIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	folderModel, err := folders.GetPortalFolderByID(ctx, portalModel, folderID)
	if err != nil {
		response.Error(w, http.StatusNotFound, err)
		return
	}

	if folderModel.PortalID != portalModel.ID {
		response.ErrorMessage(w, http.StatusForbidden, "forbidden")
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"folder": folders.PortalFolderFromModel(folderModel),
	})
}

func FolderCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)
	user, _ := middleware.UserFromContext(ctx)

	var req folders.CreateFolderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	folderModel, err := folders.FolderCreate(ctx, portalModel.ID, &req, user.ID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, map[string]any{
		"folder": folders.PortalFolderFromModel(folderModel),
	})
}

func FolderUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, _ := middleware.UserFromContext(ctx)
	portalModel, _ := middleware.PortalFromContext(ctx)

	folderIDStr := pat.Param(r, "folder_id")
	folderID, err := strconv.ParseInt(folderIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	folderModel, err := folders.GetPortalFolderByID(ctx, portalModel, folderID)
	if err != nil {
		response.Error(w, http.StatusNotFound, err)
		return
	}

	var req folders.UpdateFolderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	updatedFolder, err := folders.FolderUpdate(ctx, folderModel, &req, user.ID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"folder": folders.PortalFolderFromModel(updatedFolder),
	})
}

func FolderDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	folderIDStr := pat.Param(r, "folder_id")
	folderID, err := strconv.ParseInt(folderIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	folderModel, err := folders.GetPortalFolderByID(ctx, portalModel, folderID)
	if err != nil {
		response.Error(w, http.StatusNotFound, err)
		return
	}

	err = folders.FolderDelete(ctx, folderModel)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
