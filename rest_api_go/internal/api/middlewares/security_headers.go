package middlewares

import (
	"fmt"
	"net/http"
)

func SecurityHeaders(next http.Handler) http.Handler {
	fmt.Println("Security Headers Middleware...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Security Headers Middleware being returned...")
		w.Header().Set("X-DNS-Prefetch-Control", "off")                                             // Disable DNS prefetching, mechanism that allows browsers to resolve domain names while a user is browsing a web page
		w.Header().Set("X-Content-Type-Options", "nosniff")                                         // Prevent MIME type sniffing, which can lead to security vulnerabilities
		w.Header().Set("X-XSS-Protection", "1; mode=block")                                         // Enables the cross-site scripting (XSS) filter built into most browsers, preventing pages from loading when they detect reflected cross-site scripting attacks
		w.Header().Set("X-Frame-Options", "DENY")                                                   // Prevents the page from being displayed in a frame or iframe, mitigating clickjacking attack
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload") // Enforces secure (HTTPS) connections to the server, preventing websites from being accessed over an insecure HTTP connection like Man-in-the-Middle (MITM) attacks
		w.Header().Set("Content-Security-Policy", "default-src 'self'")                             // Defines a policy that allows resources to be loaded only from the same origin, mitigating cross-site scripting (XSS) and data injection attacks
		w.Header().Set("Referrer-Policy", "no-referrer")                                            // Controls the amount of referrer information that is passed when navigating from one page to another, enhancing privacy by not sending the referrer header
		w.Header().Set("X-Powered-By", "Go")                                                        // Indicates that the server is powered by Go, which can be useful for debugging but may also expose information about the server's technology stack
		w.Header().Set("X-Permitted-Cross-Domain-Policies", "none")                                 // Prevents Adobe Flash and Adobe Acrobat from loading content from the site, enhancing security by preventing cross-domain policy files from being used to access resources
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")           // Prevents caching of the response, ensuring that sensitive information is not stored in the browser cache
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")                               // Restricts the resources to be loaded only from the same origin, enhancing security by preventing cross-origin resource sharing (CORS) vulnerabilities
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")                              // Ensures that the page can only embed resources from the same origin or resources that explicitly allow embedding, enhancing security by preventing cross-origin resource sharing (CORS) vulnerabilities
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")                                 // Ensures that the page can only interact with other pages from the same origin, enhancing security by preventing cross-origin interactions
		w.Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")            // Controls which features can be used by the page, enhancing privacy and security by disabling potentially sensitive features like geolocation, microphone, and camera access

		next.ServeHTTP(w, r)
		fmt.Println("Security Headers ends")

	})

}

// Basic Middlware SKeleton for security headers
// func securitHeaders(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		})
// }
