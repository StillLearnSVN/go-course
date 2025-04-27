package main

import "fmt"

var middleName = "Job"

func main() {
	var age int = 20
	var name string = "John"
	var name1 = "Smith"

	// count := 10
	// lastName := "Cane"
	// middleName := "Doe"

	fmt.Println(middleName)

	fmt.Println("My name is " + name + " " + name1)
	fmt.Println("my age is ", + age)

	printname()
}

func printname() {
	firstName := "Samuel"
	fmt.Println(firstName)
}

// In conlusion, variables in go provide a flexible and powerful way to store and manipulate data. 
// They can be declared using the var keyword, and their types can be inferred or explicitly defined. 
// Go's type system allows for strong typing, which helps catch errors at compile time. 
// Additionally, Go supports various data types, including basic types like int, float64, string, and bool, 
// as well as more complex types like arrays, slices, maps, structs, interfaces, channels, functions, pointers, runes, and bytes. 
// Understanding how to use variables effectively is crucial for writing efficient and maintainable Go code.