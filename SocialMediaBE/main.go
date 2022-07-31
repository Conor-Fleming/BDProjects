package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
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
	//mux.HandleFunc("/", testHandler)
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
	case http.MethodPost:
		// call POST handler
		apiCfig.handlerCreateUser(w, r)
	case http.MethodPut:
		// call PUT handler
	case http.MethodDelete:
		// call DELETE handler
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
