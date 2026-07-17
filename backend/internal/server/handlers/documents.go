package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RhykerWells/RK/backend/internal/app/documents"
	serverErrors "github.com/RhykerWells/RK/backend/internal/server/errors"
	"github.com/RhykerWells/RK/backend/internal/server/middleware"
	"github.com/RhykerWells/RK/backend/internal/server/response"
	"goji.io/v3/pat"
)

func Document(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	docIDStr := pat.Param(r, "document_id")
	docID, err := strconv.ParseInt(docIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	documentModel, err := documents.GetDocumentByID(ctx, docID)
	if err != nil {
		response.Error(w, http.StatusNotFound, err)
		return
	}

	if documentModel.PortalID != portalModel.ID {
		response.ErrorMessage(w, http.StatusForbidden, "forbidden")
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"document": documents.PortalDocumentFromModel(documentModel),
	})
}

func DocumentCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)
	user, _ := middleware.UserFromContext(ctx)

	var req documents.CreateDocumentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	docModel, err := documents.DocumentCreate(ctx, &req, portalModel, user)
	if err != nil {
		status := http.StatusInternalServerError
		switch err {
		case serverErrors.ErrNewDocumentInvalidLocation, serverErrors.ErrInvalidFolderID, serverErrors.ErrUserNotFound, serverErrors.ErrDocumentTitleRequired:
			status = http.StatusBadRequest
		}
		response.Error(w, status, err)
		return
	}

	response.JSON(w, http.StatusCreated, map[string]any{
		"document": documents.PortalDocumentFromModel(docModel),
	})
}

func DocumentUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)
	user, _ := middleware.UserFromContext(ctx)

	docIDStr := pat.Param(r, "document_id")
	docID, err := strconv.ParseInt(docIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	docModel, err := documents.GetDocumentByID(ctx, docID)
	if err != nil {
		response.Error(w, http.StatusNotFound, err)
		return
	}

	if docModel.PortalID != portalModel.ID {
		response.ErrorMessage(w, http.StatusForbidden, "forbidden")
		return
	}

	var req documents.UpdateDocumentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	updatedDoc, err := documents.DocumentUpdate(ctx, docModel, &req, user.ID)
	if err != nil {
		status := http.StatusInternalServerError
		switch err {
		case serverErrors.ErrUpdateDocumentInvalidFolder:
			status = http.StatusBadRequest
		}
		response.Error(w, status, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"document": documents.PortalDocumentFromModel(updatedDoc),
	})
}

func DocumentDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	portalModel, _ := middleware.PortalFromContext(ctx)

	docIDStr := pat.Param(r, "document_id")
	docID, err := strconv.ParseInt(docIDStr, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	docModel, err := documents.GetDocumentByID(ctx, docID)
	if err != nil {
		response.Error(w, http.StatusNotFound, err)
		return
	}

	if docModel.PortalID != portalModel.ID {
		response.ErrorMessage(w, http.StatusForbidden, "forbidden")
		return
	}

	err = documents.DocumentDelete(ctx, docModel)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
