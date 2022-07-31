package main

import (
	"errors"
	"net/http"
)

func (apiCfig apiConfig) endpointPostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiCfig.handlerGetPost(w, r)
	case http.MethodPost:
		apiCfig.handlerCreatePost(w, r)
	case http.MethodDelete:
		apiCfig.handlerDeletePost(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}
