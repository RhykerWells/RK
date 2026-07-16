package server

import (
	"github.com/RhykerWells/RK/backend/internal/server/handlers"
	"goji.io/v3"
	"goji.io/v3/pat"
)

func registerUserRoutes(api *goji.Mux) {
	// Self User Endpoints
	api.HandleFunc(pat.Get(EndpointMe), handlers.Me)
	api.HandleFunc(pat.Post(EndpointMeAPI), handlers.IssueAPIToken)
	api.HandleFunc(pat.Delete(EndpointMeAPI), handlers.RevokeAPIToken)

	// User Retrieval Endpoints
	api.HandleFunc(pat.Get(EndpointUser), handlers.User)

	// User Management Endpoints
	api.HandleFunc(pat.Post(EndpointUser), handlers.UserCreate)
	api.HandleFunc(pat.Delete(EndpointUser), handlers.UserDelete)
}
