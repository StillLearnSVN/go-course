package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello root route!"))
	fmt.Println("Hello root route!")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on teachers route!"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method on teachers route!"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on teachers route!"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on teachers route!"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on teachers route!"))
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on students route!"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method on students route!"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on students route!"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on students route!"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on students route!"))
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on execs route!"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method on execs route!"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on execs route!"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on execs route!"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on execs route!"))
	}
}

func main() {

	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// rl := mw.NewRateLimiter(5, time.Minute) // Create a new rate limiter with a limit of 10 requests per minute

	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:             true,
	// 	CheckBody:              true,
	// 	ChecBodyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:              []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux)))))) // Apply HPP middleware with options
	// secureMux := ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	secureMux := mw.SecurityHeaders(mux)

	// Create a custom server with TLS configuration
	server := &http.Server{
		Addr: port,
		// Handler:   mux, // Use default handler
		Handler:   secureMux, // Apply CORS and security headers middleware
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port: ", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}

// Middleware is a function that takes an http.Handler and returns an http.Handler
type Middleware func(http.Handler) http.Handler

func ApplyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

// it's fine to use  HTTP handle func for simpler APIs
// using a mux provides better organization and flexibility as the API grows
// using a mux make the code more readable and maintainable
