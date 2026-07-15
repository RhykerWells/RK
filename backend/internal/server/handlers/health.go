package handlers

import (
	"net/http"

	"github.com/RhykerWells/RK/backend/internal/server/response"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, (map[string]string{
		"status": "ok",
	}))
}
