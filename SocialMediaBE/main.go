package main

import (
		"fmt"
		"net/http"
		"log"
		"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", testHandler)

	const addr = "localhost:8080"

	srv := http.Server {
		Handler: mux,
		Addr: addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
	}

	fmt.Println(addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
