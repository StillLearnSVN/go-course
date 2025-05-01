package main

import "fmt"

func main() {
	// A pointer is a variable that stores the memory address of another variable

	var ptr *int
	var a int = 10
	ptr = &a // referencing

	fmt.Println(a)
	fmt.Println(ptr)  // --> printing out the memory address
	fmt.Println(*ptr) // dereferencing the pointer to print out the actual value of a

	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int){
	*ptr++
}
