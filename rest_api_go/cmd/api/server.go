package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
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
		// teachers/{id}
		// /teachers/123
		
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")
		fmt.Println("User ID:", userID)

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

	// w.Write([]byte("Hello teachers route!"))
	// fmt.Println("Hello teachers route!")
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

	w.Write([]byte("Hello students route!"))
	fmt.Println("Hello students route!")
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

	w.Write([]byte("Hello execs route!"))
	fmt.Println("Hello execs route!")
}

func main() {

	port := ":3000"

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/teachers/", teachersHandler)

	http.HandleFunc("/students/", studentsHandler)

	http.HandleFunc("/execs/", execsHandler)

	fmt.Println("Server is running on port:", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
