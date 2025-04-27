package main

func main() {
	// Initialize a slice of integers
	// numbers := []int{1, 2, 3, 4, 5}

	// // Iterate over the slice using a for loop
	// for i := 0; i < len(numbers); i++ { 
	// 	// or we can use range for i := range numbers { --> modernized using range over int
	// 	println(numbers[i])
	// }

	// nums := []int{1, 2, 3, 4, 5}
	// for index, value := range nums {
	// 	println("Index:", index, "Value:", value)
	// }

	// for i:= range 10 {
	// 	if i%2 == 0 {
	// 		println(i, "is even")
	// 	} else {
	// 		println(i, "is odd")
	// 	}
	// }

	// // Now lets move to example of for loop with break and continue
	// for i := range 10 {
	// 	if i == 5 {
	// 		break // Exit the loop when i is 5
	// 	}
	// 	if i%2 == 0 {
	// 		continue // Skip even numbers
	// 	}
	// 	println(i) // This will only print odd numbers less than 5
	// }


	// Outer loop
	// for i:= range 5 {
	// 	// Inner loop
	// 	for j:=1; j<=5-i; j++ {
	// 		print(" ")
	// 	}
	// 	// This will print the stars, the number of stars is 2*i-1
	// 	// The number of spaces is 5-i, so that the stars are centered
	// 	for j:=1; j<=2*i-1; j++ {
	// 		print("*")
	// 	}
	// 	println() // Move to the next line after each row
	// }

	for i:= range 10 {
		println(10 - i)
		//also we can do the same using the following
		i++
		println(i)
	}
	println("done")
}