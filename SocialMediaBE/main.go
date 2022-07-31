package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Conor-Fleming/SocialMediaBE/internal/database"
)

type errorBody struct {
	Error string `json:"error"`
}

type apiConfig struct {
	dbClient database.Client
}

func main() {
	mux := http.NewServeMux()
	dbClient := database.NewClient("db.json")
	err := dbClient.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}

	apiCfig := apiConfig{
		dbClient: dbClient,
	}

	mux.HandleFunc("/users", apiCfig.endpointUsersHandler)
	mux.HandleFunc("/users/", apiCfig.endpointUsersHandler)
	mux.HandleFunc("/users/EMAIL", apiCfig.endpointUsersHandler)
	//mux.HandleFunc("/err", testErrorHandler)

	const addr = "localhost:8080"

	srv := http.Server{
		Handler:      mux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println(addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

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

/*
func testHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(200)
	//w.Write([]byte("{}"))
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}*/

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling", err)
			w.WriteHeader(500)
			response, _ := json.Marshal(errorBody{
				Error: "error marshalling",
			})
			w.Write(response)
			return
		}
		w.WriteHeader(code)
		w.Write(response)
		//w.WriteHeader(code)
	}
}

/*
func testErrorHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, errors.New("server error"))
}*/

func respondWithError(w http.ResponseWriter, code int, err error) {
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}
