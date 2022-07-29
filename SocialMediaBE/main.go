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
	Error string `json: "error"`
}

func main() {
	newDB := database.NewClient("./db.json")

	mux := http.NewServeMux()
	mux.HandleFunc("/", testHandler)
	mux.HandleFunc("/err", testErrorHandler)

	const addr = "localhost:8080"

	srv := http.Server{
		Handler:      mux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println(addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(200)
	//w.Write([]byte("{}"))
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
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
	w.Write(response)
	w.WriteHeader(code)
}

func testErrorHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, errors.New("server error"))
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}
