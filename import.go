package main

import (
	"fmt"
	foo "net/http"
)

func main() {

	fmt.Println("Hello, Go Standard Library")
	// Define a URL to fetch
	url := "https://api.github.com"

	// Make a GET request to the URL
	resp, err := foo.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Ensure the response body is closed after reading
	defer resp.Body.Close()

	// Check if the response status code is 200 OK
	if resp.StatusCode == foo.StatusOK {
		fmt.Println("Response Status:", resp.Status)
	} else {
		fmt.Println("Error: Received non-200 response status:", resp.Status)
	}

	// Tree Shaking explanation:
	// The Go compiler automatically removes unused code during the build process.
	// This means that if you import a package but don't use any of its functions or types,
	// the compiler will not include that package in the final binary.
	// This is known as "tree shaking" and helps keep the binary size small.
	// In this example, we imported the "net/http" package and used its Get function,
	// but if we had not used it, the compiler would have removed the import from the final binary.
	// This is a powerful feature of Go that helps developers create efficient and optimized binaries.
	// The Go compiler is smart enough to only include the code that is actually used in the final binary. 
	// Tree shaking efficiently trims unnecessary parts during the build process.

	// In what term of Go Import adopted tree shaking?
	// The term "tree shaking" is not commonly used in the context of Go imports.
	// Instead, Go uses a concept called "dead code elimination" to refer to the process of removing unused code.
	// This process is part of the Go compiler's optimization phase, where it analyzes the code and removes any functions, variables, or packages that are not referenced in the program.
	// This helps to keep the final binary size small and improves performance.
	// The Go compiler performs dead code elimination during the build process, which is similar to tree shaking in other programming languages.
	// In summary, while the term "tree shaking" is not commonly used in Go, the concept of removing unused code is an integral part of the Go compiler's optimization process.
}