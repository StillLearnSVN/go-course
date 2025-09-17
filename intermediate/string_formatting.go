package main

import "fmt"

func main() {
	num := 42
	fmt.Printf("%05d\n", num) // Output: 00042, %05 means pad with zeros to a width of 5

	message := "Hello"
	fmt.Printf("|%-10s!|\n", message) // Output: "Hello     !", %-10 means left-align in a width of 10
	fmt.Printf("|%10s!|\n", message)  // Output: "     Hello!", %10 means right-align in a width of 10

	msg1 := "Go"
	msg2 := `Go\nLang`
	fmt.Println(msg1) // Output: Go
	fmt.Println(msg2) // Output: Go\nLang

	// Using backticks for raw string literals
	rawString := `This is a raw string literal.
It can span multiple lines.
Special characters like '\n' are not interpreted.`
	fmt.Println(rawString)

	// While using backticks, you cannot include backticks inside the string directly.
	// and for using double quotes inside backticks, you can do it directly.
	// Double quotes is used for interpreted string literals.


}
