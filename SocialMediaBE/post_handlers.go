package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (apiCfig apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserEmail string `json:"userEmail"`
		Text      string `json:"text"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	result, err := apiCfig.dbClient.CreatePost(params.UserEmail, params.Text)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, result)
}

func (apiCfig apiConfig) handlerGetPost(w http.ResponseWriter, r *http.Request) {
	resEmail := strings.TrimPrefix(r.URL.Path, "/posts/")
	result, err := apiCfig.dbClient.GetPosts(resEmail)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, result)
}

func (apiCfig apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	resUuid := strings.TrimPrefix(r.URL.Path, "/posts/")
	err := apiCfig.dbClient.DeletePost(resUuid)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, struct{}{})
}
