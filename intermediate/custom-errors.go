package main

import (
	"errors"
	"fmt"
)

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Program continues...")
}

type customError struct {
	code int
	msg  string
	er   error
}

// Error returns the error message.. Implementing Error() method of error interface
func (ce *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v", ce.code, ce.msg, ce.er)
}

// function that returns a custom error
// func doSomething() error {
// 	return &customError{code: 404, msg: "Resource not found"}
// }

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{code: 500, msg: "Internal Server Error", er: err}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("A standard error occurred")
}
