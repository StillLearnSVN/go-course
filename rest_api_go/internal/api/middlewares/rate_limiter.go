package middlewares

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	visitors  map[string]int
	limit     int
	resetTime time.Duration
}

func NewRateLimiter(limit int, resetTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors:  make(map[string]int),
		limit:     limit,
		resetTime: resetTime,
	}
	// Start a goroutine to reset the visitor counts after the specified reset time
	go rl.resetVisitorCount()
	return rl
}

func (rl *RateLimiter) resetVisitorCount() {
	for {
		time.Sleep(rl.resetTime)           // Wait for the reset time
		rl.mu.Lock()                       // Lock to safely modify the visitors map
		rl.visitors = make(map[string]int) // Reset the visitor counts
		rl.mu.Unlock()                     // Unlock after modification
	}
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	fmt.Println("Rate Limiter Middleware...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Rate Limiter Middleware being returned...")
		rl.mu.Lock()         // Lock to safely access the visitors map
		defer rl.mu.Unlock() // Ensure the mutex is unlocked after processing

		visitorIp := r.RemoteAddr // Get the visitor's visitorIp address
		if count, exists := rl.visitors[visitorIp]; exists && count >= rl.limit {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests) // Return 429 status if limit exceeded
			return
		}

		// Increment the visitor count
		rl.visitors[visitorIp]++
		// fmt.Println("Visitor IP:", visitorIp, "Count:", rl.visitors[visitorIp])

		if rl.visitors[visitorIp] > rl.limit {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests) // Return 429 status if limit exceeded
			return
		}
		next.ServeHTTP(w, r) // Call the next handler in the chain
		fmt.Println("Rate Limiter ends...")

	})
}
