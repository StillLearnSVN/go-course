package main

import "fmt"

type GoogleEmployee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// This is used for exported identifiers (public)
	// Example: MyVariable, GetUserName, CalculateArea, UserInfo or datatype like struct, interface

	// snake_case
	// This is commonly used for naming variables, constants, and file names
	// E.g : my_variable, user_name, calculate_area, user_info

	//UPPERCASE
	// This is used for constants
	// Example: MAX_VALUE, PI, DEFAULT_TIMEOUT

	// mixedCase
	// This is used for local variables and function names
	// Example: myVariable, getUserName, calculateArea, userInfo

	// camelCase
	// This is used for local variables and function names
	// Example: myVariable, getUserName, calculateArea, userInfo

	const MAX_ENTRIES = 10

	var employeeID = 12345
	fmt.Println("Employee ID:", employeeID)
}
