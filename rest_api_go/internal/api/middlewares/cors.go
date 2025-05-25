package middlewares

import (
	"fmt"
	"net/http"
)

var allowedOrigins = []string{
	"https://localhost:3000",
	"https://example.com",
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Println("CORS Origin:", origin)

		if isOriginAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin) // Sets the Access-Control-Allow-Origin header to the origin of the request, allowing the browser to access the resource from that origin
		} else {
			http.Error(w, "CORS not allowed for this origin", http.StatusForbidden)
			return
		}
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Specifies which headers can be used in the actual request, allowing the server to control which headers are allowed in cross-origin requests
		w.Header().Set("Access-Control-Expose-Headers", "Authorization") // Specifies which headers can be exposed to the browser, allowing the server to control which headers are accessible in cross-origin requests
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS") // Specifies which HTTP methods are allowed when accessing the resource, allowing the server to control which methods can be used in cross-origin requests
		w.Header().Set("Access-Control-Allow-Credentials", "true") // Indicates whether the browser should include credentials (cookies, HTTP authentication, and client-side SSL certificates) in cross-origin requests, allowing the server to control whether credentials are allowed in cross-origin requests
		w.Header().Set("Access-Control-Max-Age", "3600") // Specifies how long the results of a preflight request can be cached, allowing the server to control how long the browser can cache the preflight response

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isOriginAllowed(origin string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}
	return false
}