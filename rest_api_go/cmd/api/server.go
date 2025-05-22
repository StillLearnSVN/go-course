package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  string    `json:"age"`
	City string `json:"city"`
}

func main() {

	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello root route!"))
		fmt.Println("Hello root route!")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

			// Access the request details
			fmt.Println("Body:", r.Body)
			fmt.Println("Form:", r.Form)
			fmt.Println("Header:", r.Header)
			fmt.Println("Context:", r.Context())
			fmt.Println("Context Length:", r.ContentLength)
			fmt.Println("Host:", r.Host)
			fmt.Println("Method:", r.Method)
			fmt.Println("Protocol:", r.Proto)
			fmt.Println("Remote Address:", r.RemoteAddr)
			fmt.Println("Request URI:", r.RequestURI)
			fmt.Println("TLS:", r.TLS)
			fmt.Println("URL:", r.URL)
			fmt.Println("Trailer:", r.Trailer)
			fmt.Println("Transfer Encoding:", r.TransferEncoding)
			fmt.Println("User Agent:", r.UserAgent())
			fmt.Println("Port", r.URL.Port())
			fmt.Println("URL Scheme", r.URL.Scheme)

			w.Write([]byte("Hello GET method on teachers route!"))
			fmt.Println("Hello GET method  on teachers route!")
		case http.MethodPost:
			// Parse form data (necessary for x-www-form-urlencoded)
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form data", http.StatusBadRequest)
				return
			}

			fmt.Println("Form data:", r.Form)

			// Prepare response data
			response := make(map[string]interface{})
			for key, value := range r.Form {
				response[key] = value[0] // Take the first value for each key
			}

			fmt.Println("Response data:", response)

			// RAW Body
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusBadRequest)
				return
			}
			defer r.Body.Close() // Close the body after reading
			fmt.Println("Raw body:", body)
			fmt.Println("Raw body:", string(body))

			// If you expect JSON data, you can unmarshal it into a struct
			var userInstance user
			err = json.Unmarshal(body, &userInstance)

			if err != nil {
				http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
				return
			}

			// Print the unmarshalled data
			fmt.Println("Unmarshalled JSON into a user struct:", userInstance)
			fmt.Println("Received username as:", userInstance.Name)

			// Prepare response data
			response1 := make(map[string]interface{})
			for key, value := range r.Form {
				response[key] = value[0] // Take the first value for each key
			}

			err = json.Unmarshal(body, &response1)

			if err != nil {
				http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
				return
			}

			fmt.Println("Unmarshalled JSON into a map:", response1)

			// Access the request details
			fmt.Println("Body:", r.Body)
			fmt.Println("Form:", r.Form)
			fmt.Println("Header:", r.Header)
			fmt.Println("Context:", r.Context())
			fmt.Println("Context Length:", r.ContentLength)
			fmt.Println("Host:", r.Host)
			fmt.Println("Method:", r.Method)
			fmt.Println("Protocol:", r.Proto)
			fmt.Println("Remote Address:", r.RemoteAddr)
			fmt.Println("Request URI:", r.RequestURI)
			fmt.Println("TLS:", r.TLS)
			fmt.Println("URL:", r.URL)
			fmt.Println("Trailer:", r.Trailer)
			fmt.Println("Transfer Encoding:", r.TransferEncoding)
			fmt.Println("User Agent:", r.UserAgent())
			fmt.Println("Port", r.URL.Port())
			fmt.Println("URL Scheme", r.URL.Scheme)

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
		// if r.Method != http.MethodGet {
		// 	w.Write([]byte("Hello GET method on teachers route!"))
		// 	fmt.Println("Hello GET method on teachers route!")
		// 	return
		// }

		w.Write([]byte("Hello teachers route!"))
		fmt.Println("Hello teachers route!")
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {

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
	})

	fmt.Println("Server is running on port:", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
