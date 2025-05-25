package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received response in ResponseTime")
		start := time.Now() // Record the start time of the request

		// Create a custom ResponseWriter to capture the status code
		wrappedWriter := &responseTimeWriter{
			ResponseWriter: w,
			status:          http.StatusOK, // Default status code
		}

		
		// Calculate the response time
		duration := time.Since(start)
		wrappedWriter.Header().Set("X-Response-Time", duration.String()) // Set a custom header to indicate that the response time is being tracked
		
		next.ServeHTTP(wrappedWriter, r) // Call the next handler in the chain
		// Log the response time and status code
		fmt.Println("Method:", r.Method, "URL:", r.URL.Path, "Status Code:", wrappedWriter.status, "Response Time:", duration)
		fmt.Println("Sent Response from Response Time Middleware")
	})
}

// ResponseWriter
type responseTimeWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseTimeWriter) WriteHeader(status int) {
	w.status = status // Capture the status code
	w.ResponseWriter.WriteHeader(status) // Call the original WriteHeader method
}