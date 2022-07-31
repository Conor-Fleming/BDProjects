package main

import (
	"errors"
	"net/http"
)

func (apiCfig apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// call GET handler
		apiCfig.handlerGetUser(w, r)
	case http.MethodPost:
		// call POST handler
		apiCfig.handlerCreateUser(w, r)
	case http.MethodPut:
		// call Update handler
		apiCfig.handlerUpdateUser(w, r)
	case http.MethodDelete:
		//call Delete handler
		apiCfig.handlerDeleteUser(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}
