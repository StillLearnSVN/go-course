package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"strconv"
	"strings"
	// "sync"
)

type Teacher struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Class     string `json:"class"`
	Subject   string `json:"subject"`
}

var (
	teachers = make(map[int]Teacher)
	// mutex    = &sync.Mutex{}
	nextID = 1
)

// Initializes the teachers map with some sample data
func init() {
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Algebra",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10A",
		Subject:   "Biology",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "11A",
		Subject:   "Computer Programming",
	}
}

func getTeacherHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	fmt.Println(idStr)

	if idStr == "" {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teacherList := make([]Teacher, 0, len(teachers))

		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
				teacherList = append(teacherList, teacher)
			}

		}
		response := struct {
			Status string    `json:"status"`
			Count  int       `json:"count"`
			Data   []Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teacherList),
			Data:   teacherList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	// Handle Path Parameters
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return 
	}

	teacher, exists := teachers [id]
	if !exists {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(teacher)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello root route!"))
	fmt.Println("Hello root route!")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// call get method handler
		getTeacherHandler(w, r)
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
