package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func loadClientCAs() *x509.CertPool {
	clientCAs := x509.NewCertPool()
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatalln("Error reading CA certificate:", err)
	}
	clientCAs.AppendCertsFromPEM(caCert)
	return clientCAs
}

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request){
		// Log the request details
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling incoming orders!")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request){
		// Log the request details
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling users!")
	})

	port := 3000

	// Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ClientAuth: tls.RequireAndVerifyClientCert, // Require client certificate(enforce mutual TLS)
		ClientCAs: loadClientCAs(),

		// In mutual TLS, both server and client authenticate each other
		// and server verifies the client certificate ensuring only authorized client can connect
	}

	// Create a custom server
	server := &http.Server{
		Addr: 	fmt.Sprintf(":%d", port),
		Handler: nil,
		TLSConfig: tlsConfig,	
	}

	// Enable HTTP/2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Starting server on port", port)
	// HTTP/2 server with TLS
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}

	

	// HTTP 1.1 server without TLS
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatalln("Error starting server:", err)
	// }

	
}

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("HTTP Version:", httpVersion)

	if r.TLS != nil {
		tlsVersion := r.TLS.Version
		fmt.Println("TLS Version:", getTLSVersionName(tlsVersion))
	} else {
		fmt.Println("No TLS version")
	}
	fmt.Println("Request Method:", r.Method)
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return fmt.Sprintf("Unknown TLS version: %d", version)
	}
}
