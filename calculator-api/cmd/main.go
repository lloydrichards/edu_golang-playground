package main

import (
	"log"
	"net/http"
	"time"

	"github.com/lloydrichards/calculator_api/internal/handlers"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/add", handlers.HandleAdd)
	router.HandleFunc("/subtract", handlers.HandleSubtract)
	router.HandleFunc("/multiply", handlers.HandleMultiply)
	router.HandleFunc("/divide", handlers.HandleDivide)
	router.HandleFunc("/sum", handlers.HandleSum)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware(router),
	}

	log.Println("Starting the server on http://localhost:8080")
	server.ListenAndServe()
}
