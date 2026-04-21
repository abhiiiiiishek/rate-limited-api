package handler

import (
	"encoding/json"
	"net/http"

	"rate-limited-api/internal/limiter"
	"rate-limited-api/internal/model"
)

type Handler struct {
	limiter *limiter.RateLimiter
}

func NewHandler(l *limiter.RateLimiter) *Handler {
	return &Handler{limiter: l}
}

func (h *Handler) RequestHandler(w http.ResponseWriter, r *http.Request) {
	var req model.Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	allowed := h.limiter.Allow(req.UserID)
	if !allowed {
		w.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "rate limit exceeded",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status": "request accepted",
	})
}

func (h *Handler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	stats := h.limiter.Stats()
	json.NewEncoder(w).Encode(stats)
}
