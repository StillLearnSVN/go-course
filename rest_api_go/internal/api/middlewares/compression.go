package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	fmt.Println("Compression Middleware..")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Compression Mifdleware being returned..")
		// Check if the request accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r) // If not, just call the next handler
		}

		// Set the response header to indicate gzip encoding
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w) // Create a new gzip writer
		defer gz.Close()        // Ensure the gzip writer is closed after use

		// Wrap the original ResponseWriter with our gzipResponseWriter
		w = &gzipResponseWriter{
			ResponseWriter: w,
			gzWriter:       gz, // Assign the gzip writer to the custom ResponseWriter
		}

		next.ServeHTTP(w, r) // Call the next handler in the chain

		fmt.Println("Sent Response from Compression Middleware")
	})
}

// gzipResponseWriter is a custom ResponseWriter that wraps the original ResponseWriter
type gzipResponseWriter struct {
	http.ResponseWriter
	gzWriter *gzip.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	if g.gzWriter == nil {
		return g.ResponseWriter.Write(b) // If gzWriter is nil, write directly to the ResponseWriter
	}
	return g.gzWriter.Write(b) // Write compressed data to the gzip writer
}
