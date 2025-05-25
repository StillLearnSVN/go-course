package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

type HPPOptions struct {
	CheckQuery             bool
	CheckBody              bool
	ChecBodyForContentType string
	Whitelist              []string
}

func Hpp(options HPPOptions) func(http.Handler) http.Handler {
	fmt.Println("HPP Middleware...")
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("HPP Middleware being returned...")
			if options.CheckBody && r.Method == http.MethodPost && isCorrectContentType(r, options.ChecBodyForContentType) {
				//filter the body params
				filterBodyParams(r, options.Whitelist)
			}

			if options.CheckQuery && r.URL.Query() != nil {
				//filter the query params
				filterQueryParams(r, options.Whitelist)
			}

			next.ServeHTTP(w, r) // Call the next handler in the chain
			fmt.Println("HPP Middleware end..")

		})
	}
}

func isCorrectContentType(r *http.Request, contentType string) bool {
	return strings.Contains(r.Header.Get("Content-Type"), contentType)
}

func filterBodyParams(r *http.Request, whitelist []string) {
	err := r.ParseForm() // Parse the form data from the request body
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range r.Form {
		if len(v) > 1 {
			r.Form.Set(k, v[0]) // If there are multiple values for a key, keep only the first one
			// r.Form.Set(k, v[len(v)-1]) // Keep the last value as well
		}
		if !isWhitelisted(k, whitelist) {
			delete(r.Form, k) // If the parameter is not whitelisted, delete it from the form
		}
	}

}

func filterQueryParams(r *http.Request, whitelist []string) {
	query := r.URL.Query() // Get the query parameters from the URL

	for k, v := range query {
		if len(v) > 1 {
			// query.Set(k, v[0]) // If there are multiple values for a key, keep only the first one
			query.Set(k, v[len(v)-1]) // Keep the last value as well
		}
		if !isWhitelisted(k, whitelist) {
			query.Del(k) // If the parameter is not whitelisted, delete it from the query
		}
	}
	r.URL.RawQuery = query.Encode() // Update the URL with the modified query parameters
}

func isWhitelisted(param string, whitelist []string) bool {
	for _, v := range whitelist {
		if param == v {
			return true // If the parameter is in the whitelist, return true
		}
	}
	return false // If not found in the whitelist, return false
}
