package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/internal/repositories/sqlconnect"

	"github.com/joho/godotenv"
)


func main() {

	err := godotenv.Load()
	if err != nil {
		return
	}

	_, err = sqlconnect.ConnectDb() // Connect to the database
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	
	port := os.Getenv("API_PORT")

	cert := "cert.pem"
	key := "key.pem"

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
	router := router.MainRouter()
	secureMux := mw.SecurityHeaders(router)

	// Create a custom server with TLS configuration
	server := &http.Server{
		Addr: port,
		// Handler:   mux, // Use default handler
		Handler:   secureMux, // Apply CORS and security headers middleware
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port: ", port)

	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
