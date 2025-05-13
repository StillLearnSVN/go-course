package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Person struct
type Person struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

// Sample data
var personData = map[string]Person{
	"1": {Name: "John Doe", Age: 30},
	"2": {Name: "Jane Doe", Age: 28},
	"3": {Name: "Jack Doe", Age: 25},
}

// Handler function for the endpoint
func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL query parameters
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	// Check if the ID exists in the personData map
	person, exists := personData[id]
	if !exists {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the person data to JSON and write to the response
	if err := json.NewEncoder(w).Encode(person); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	// Define the port
	port := 8080

	// Print the confirmation message
	fmt.Printf("Server started on port %d\n", port)

	// Set up the endpoint and the handler function
	http.HandleFunc("/person", getPersonHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}