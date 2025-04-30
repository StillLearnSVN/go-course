package main

import "fmt"

// The init function in Go, provides a mechanism to initialize package-level variables and perform setup tasks.

func init() {
	fmt.Println("Initializing the package1...")
}
func init() {
	fmt.Println("Initializing the package2...")
}
func init() {
	fmt.Println("Initializing the package3...")
}

// Init functions are executed in the order they are defined in the file.
// If there are multiple init functions in the same package, they are executed in the order they appear in the file.
// If there are multiple init functions in different packages, they are executed in the order of package dependency.
// The main function is the entry point of the program.
// The init functions are executed before the main function.

func main() {
	fmt.Println("Main function execution...")
}
