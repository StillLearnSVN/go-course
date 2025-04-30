package basics

import "fmt"

func main() {
	message := "Hello World"

	for i, v := range message {
		// fmt.Println(i, v)
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}
	// The above code will print the index and the rune value of each character in the string "Hello World".
}