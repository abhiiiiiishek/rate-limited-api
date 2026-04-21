package main

import (
	"log"
	"net/http"
	"time"

	"rate-limited-api/internal/handler"
	"rate-limited-api/internal/limiter"
)

func main() {
	rl := limiter.NewRateLimiter(5, time.Minute)
	h := handler.NewHandler(rl)

	mux := http.NewServeMux()
	mux.HandleFunc("/request", h.RequestHandler)
	mux.HandleFunc("/stats", h.StatsHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
