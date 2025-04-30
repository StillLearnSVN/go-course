package main

import "fmt"

func main() {
	// func <name>(<parameters>) <return type> {
	// code block
	// return value
	//}

	// fmt.Println(add(1, 2))

	// greet := func() { // This is an anonymous function, it doesn't have a name
	// 	fmt.Println("Hello anonymous function")
	// }

	// greet() // This will call the anonymous function and print "Hello anonymous function"

	// operation := add

	// result := operation(3,3)

	// fmt.Println(result)

	// Passing a function as an argument
	result := applyOperation(3, 4, add)
	fmt.Println("3 + 4 = ", result)

	// Returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 = ", multiplyBy2(6))

}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)

	/*
	This function takes three arguments: two integers (x and y) and a function (operation) that itself takes two integers and returns an integer. 
	In this case, the add function is passed as the operation argument, which simply adds the two integers. 
	The applyOperation function then calls the passed-in function (add) with the provided integers (3 and 4) and returns the result. 
	This demonstrates how you can pass behavior (in the form of a function) as an argument, making the applyOperation function highly versatile.
	*/
}

// Function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}

	/*
	The createMultiplier function takes an integer factor as an argument and returns a new function. This returned function takes an integer x 
	and multiplies it by the factor. In the example, createMultiplier is called with 2, creating a function (multiplyBy2) 
	that doubles any number passed to it. When multiplyBy2(6) is called, it multiplies 6 by 2 and returns 12. 
	This pattern is useful for creating specialized functions dynamically based on input parameters.
	*/
}
