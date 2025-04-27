package main

import "math"

const (
	PI = 3.14
	E  = 2.718
	// Constants can be of any type
	// including complex numbers, strings, and booleans.
	// Complex number
	COMPLEX = 1 + 2i
	// String
	STRING = "Hello, World!"
	// Boolean
	BOOLEAN = true
)

//undefined exp func, let's define it as real function
// exp function to demonstrate usage of constants
func exp(x float64) float64 {
	// Using the real function exponential it self
	return math.Exp(x)
}

func main() {
	// Constants are immutable and cannot be changed.
	// PI = 3.14159 // This will cause a compilation error
	// E = 2.71828 // This will also cause a compilation error

	// You can use constants in expressions
	circumference := 2 * PI * 5
	println("Circumference of circle with radius 5:", circumference)

	// You can also use constants in functions
	println("Exponential of 1:", exp(1))

}

// In Conlusion, constants in Go provide a mechanism for defining immutable values that can be used throughout your code. 
// They are useful for defining fixed values that do not change, such as mathematical constants, configuration values, and more.
