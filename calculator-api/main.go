package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello world"}`))
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting the server on http://localhost:8080")
	server.ListenAndServe()
}
