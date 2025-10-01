package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"restapi/pkg/utils"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

func XSSMiddleware(next http.Handler) http.Handler {
	fmt.Println("********** Initializing XSSMiddleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Here you would add your XSS sanitization logic
		// For demonstration, we'll just print a message
		fmt.Println("XSSMiddleware: Sanitizing request")

		// Sanitize the URL path
		sanitizedPath, err := clean(r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("Original Path:", r.URL.Path)
		fmt.Println("Sanitized Path:", sanitizedPath)

		// Sanitize query parameters
		params := r.URL.Query()
		sanitizedQuery := make(map[string][]string)
		for key, values := range params {
			sanitizedKey, err := clean(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			var sanitizedValues []string
			for _, value := range values {
				sanitizedValue, err := clean(value)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				sanitizedValues = append(sanitizedValues, sanitizedValue.(string))
			}
			sanitizedQuery[sanitizedKey.(string)] = sanitizedValues
			fmt.Printf("Original Query Param: %s=%s, \nSanitized Query: %s=%s\n", key, strings.Join(values, ", "), sanitizedKey, strings.Join(sanitizedValues, ", "))
		}

		r.URL.Path = sanitizedPath.(string)
		r.URL.RawQuery = url.Values(sanitizedQuery).Encode()
		fmt.Println("Updated URL:", r.URL.String())

		// Sanitize request body
		if r.Header.Get("Content-Type") == "application/json" {
			if r.Body != nil {
				bodyBytes, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(w, utils.ErrorHandler(err, "Failed to read request body").Error(), http.StatusInternalServerError)
					return
				}

				bodyString := strings.TrimSpace(string(bodyBytes))
				fmt.Println("Original Body:", bodyString)

				// Reset the request body so it can be read again
				r.Body = io.NopCloser(bytes.NewReader([]byte(bodyString)))

				if len(bodyString) > 0 {
					var inputData interface{}
					err := json.NewDecoder(bytes.NewReader([]byte(bodyString))).Decode(&inputData)
					if err != nil {
						http.Error(w, utils.ErrorHandler(err, "Invalid JSON in request body").Error(), http.StatusBadRequest)
						return
					}
					fmt.Println("Original JSON Body:", inputData)

					// Sanitize the JSON body
					sanitizedData, err := clean(inputData)
					if err != nil {
						http.Error(w, "Failed to sanitize request body", http.StatusInternalServerError)
						return
					}
					fmt.Println("Sanitized JSON Body:", sanitizedData)

					// Marshal the sanitized data back to JSON
					sanitizedBody, err := json.Marshal(sanitizedData)
					if err != nil {
						http.Error(w, utils.ErrorHandler(err, "Error sanitizing body").Error(), http.StatusInternalServerError)
						return
					}

					r.Body = io.NopCloser(bytes.NewReader(sanitizedBody))
					fmt.Println("Sanitized Body:", string(sanitizedBody))

				} else {
					log.Println("Request body is empty after trimming whitespace, skipping body sanitization")
				}

			} else {
				log.Println("Request body is nil, skipping body sanitization")
			}
		} else if r.Header.Get("Content-Type") != "" {
			log.Printf("Received request with unsupported content-type: %s. Expected application/json", r.Header.Get("Content-Type"))
			http.Error(w, "Unsupported content-type. Only application/json is supported.", http.StatusUnsupportedMediaType)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
		fmt.Println("XSSMiddleware: Finished processing request")
	})
}

// Clean sanitize input to prevent XSS attacks
func clean(data interface{}) (interface{}, error) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, val := range v {
			v[key] = sanitizeValue(val)
		}
		return v, nil

	case []interface{}:
		for i, val := range v {
			v[i] = sanitizeValue(val)
		}
		return v, nil
	case string:
		return sanitizeString(v), nil
	default:
		return nil, utils.ErrorHandler(fmt.Errorf("unsupported data type: %T", data), fmt.Sprintf("data: %v", data))
	}
}

func sanitizeValue(data interface{}) interface{} {
	switch v := data.(type) {
	case string:
		return sanitizeString(v)
	case map[string]interface{}:
		for key, val := range v {
			v[key] = sanitizeValue(val)
		}
		return v

	case []interface{}:
		for i, val := range v {
			v[i] = sanitizeValue(val)
		}
		return v
	default:
		return v // Return v as it is for unsupported types
	}
}

func sanitizeString(value string) string {
	return bluemonday.UGCPolicy().Sanitize(value)
}
