package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//Handler functions Create, Update, Get , and Delete

//handlerCreateUser takes a http response writer and a request with a body containing the info needed to create a user
//the function take the body from the request and uses that data to call CreateUser in the Database packaage.
//if all goes well the function will respond with JSON and display the new user.
func (apiCfig apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	result, err := apiCfig.dbClient.CreateUser(params.Email, params.Password, params.Name, params.Age)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, result)
}

func (apiCfig apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	resEmail := strings.TrimPrefix(r.URL.Path, "/users/")
	fmt.Println(resEmail, r.URL.Path)
	err := apiCfig.dbClient.DeleteUser(resEmail)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (apiCfig apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	resEmail := strings.TrimPrefix(r.URL.Path, "/users/")
	result, err := apiCfig.dbClient.GetUser(resEmail)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, result)
}

func (apiCfig apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	resEmail := strings.TrimPrefix(r.URL.Path, "/users/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	result, err := apiCfig.dbClient.UpdateUser(resEmail, params.Password, params.Name, params.Age)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, result)
}
