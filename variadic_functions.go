package main

import "fmt"

func main() {
	// Variadic function in go allows you to pass a variable number of arguments to a function.
	// Allow you to create functions that can accept a variable number of arguments.

	// ... Ellipsis (...) is used to define a variadic function.

	// func functionName(param1 type1, param2 type2, ...paramN typeN) returnType {
		// function body
	// }

	// fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3)) // Output: Sum of 1, 2, 3: 6
	sequence, total := sum(1, 1, 2, 3)
	fmt.Println("Sequence: ",sequence, "Total: ", total) // Output: Sequence:  1 Total:  6

	sequence2, total2 := sum(2, 11, 42, 23)
	fmt.Println("Sequence: ",sequence2, "Total: ", total2) // Output: Sequence:  2 Total:  76

	numbers := []int{1, 2, 3, 4, 5}

	// You can also pass a slice to a variadic function
	// You can use the ... operator to pass a slice to a variadic function
	seq3, total3 := sum(3, numbers...)
	fmt.Println("Sequence: ",seq3, "Total: ", total3) // Output: Sequence:  3 Total:  15
}

func sum(sequence int, nums ...int) (int, int ){
	// numbers is a slice of integers
	// You can pass any number of integers to this function
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}

// In conclusion, variadic functions in Go are a powerful feature that allows you to create flexible and reusable functions that can accept a variable number of arguments.
// Also provide a flexible way to define functions that can accept a variable number of arguments of a specific type.