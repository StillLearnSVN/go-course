package main

import "fmt"

func main() {
	// In Go, the `defer` statement is used to ensure that a function call is performed later 
	// in the program's execution, usually for purposes of cleanup.
	process(10)

}

func process(i int) {
	// The deferred function calls are pushed onto a stack. When the surrounding function returns,
	// the deferred calls are executed in last-in-first-out order.
	// This means that the last deferred function call will be executed first, 
	// and the first deferred function call will be executed last.
	defer fmt.Println("Deffered i value:", i)
	defer fmt.Println("First Deferred call executed")
	defer fmt.Println("Second Deferred call executed")
	defer fmt.Println("Third Deferred call executed")
	i++
	fmt.Println("Normal call executed")
	fmt.Println("Value of i is:", i)
}
