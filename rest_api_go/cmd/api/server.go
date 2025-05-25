package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello root route!"))
	fmt.Println("Hello root route!")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on teachers route!"))
		// fmt.Println("Hello GET method  on teachers route!")
	case http.MethodPost:
		w.Write([]byte("Hello POST method on teachers route!"))
		fmt.Println("Hello POST method on teachers route!")
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on teachers route!"))
		fmt.Println("Hello PUT method on teachers route!")
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on teachers route!"))
		fmt.Println("Hello DELETE method on teachers route!")
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on teachers route!"))
		fmt.Println("Hello PATCH method on teachers route!")
		return
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on students route!"))
		fmt.Println("Hello GET method  on students route!")
	case http.MethodPost:
		w.Write([]byte("Hello POST method on students route!"))
		fmt.Println("Hello POST method on students route!")
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on students route!"))
		fmt.Println("Hello PUT method on students route!")
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on students route!"))
		fmt.Println("Hello DELETE method on students route!")
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on students route!"))
		fmt.Println("Hello PATCH method on students route!")
		return
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on execs route!"))
		fmt.Println("Hello GET method  on execs route!")
	case http.MethodPost:
		w.Write([]byte("Hello POST method on execs route!"))
		fmt.Println("Hello POST method on execs route!")
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on execs route!"))
		fmt.Println("Hello PUT method on execs route!")
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on execs route!"))
		fmt.Println("Hello DELETE method on execs route!")
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on execs route!"))
		fmt.Println("Hello PATCH method on execs route!")
		return
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

	// Create a custom server with TLS configuration
	server := &http.Server{
		Addr: port,
		// Handler:   mux, // Use default handler
		Handler:   mw.Cors(mw.SecurityHeaders(mux)), // Apply CORS and security headers middleware
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port: ", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}

// it's fine to use  HTTP handle func for simpler APIs
// using a mux provides better organization and flexibility as the API grows
// using a mux make the code more readable and maintainable
