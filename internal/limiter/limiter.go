package limiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu           sync.Mutex
	userRequests map[string][]time.Time
	limit        int
	window       time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		userRequests: make(map[string][]time.Time),
		limit:        limit,
		window:       window,
	}
}

func (r *RateLimiter) Allow(userID string) bool {
	now := time.Now()

	r.mu.Lock()
	defer r.mu.Unlock()

	timestamps := r.userRequests[userID]

	valid := make([]time.Time, 0, len(timestamps))
	for _, t := range timestamps {
		if now.Sub(t) < r.window {
			valid = append(valid, t)
		}
	}

	if len(valid) >= r.limit {
		r.userRequests[userID] = valid
		return false
	}

	valid = append(valid, now)
	r.userRequests[userID] = valid

	return true
}

func (r *RateLimiter) Stats() map[string]int {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make(map[string]int)
	now := time.Now()

	for user, timestamps := range r.userRequests {
		count := 0
		for _, t := range timestamps {
			if now.Sub(t) < r.window {
				count++
			}
		}
		if count > 0 {
			result[user] = count
		}
	}

	return result
}
