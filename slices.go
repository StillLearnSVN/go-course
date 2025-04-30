package main

import (
	"fmt"
	"slices"
)

func main() {
	// var numbers []int
	// var numbers1 = []int{1, 2, 3}

	// numbers2 := []int{9,8,7} // numbers2 is a slice of integers

	// slice := make([]int, 5) // slice of length 5

	a := [5]int{1, 2, 3, 4, 5} // array of 5 integers
	slice1 := a[1:4]           // slice of array that will take elements from index 1 to 3

	fmt.Println("slice1:", slice1) // Output: slice1: [2 3 4]

	slice1 = append(slice1, 6, 7) // append 6 to slice1

	fmt.Println("slice1 after append:", slice1) // Output: slice1 after append: [2 3 4 6 7]

	sliceCopy := make([]int, len(slice1)) // create a new slice of the same length as slice1
	copy(sliceCopy, slice1)               // copy elements from slice1 to sliceCopy

	fmt.Println("sliceCopy:", sliceCopy) // Output: sliceCopy: [2 3 4 6 7]

	var nilSlice []int                 // nil slice
	fmt.Println("nilSlice:", nilSlice) // Output: nilSlice: []

	for i, v := range slice1 {
		fmt.Printf("slice1[%d]: %d\n", i, v) // Output: slice1[0]: 2, slice1[1]: 3, slice1[2]: 4, slice1[3]: 6, slice1[4]: 7
	}

	fmt.Println("Element at index 2:", slice1[2]) // Output: Element at index 2: 4

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 and sliceCopy are equal") // Output: slice1 and sliceCopy are equal
	} else {
		fmt.Println("slice1 and sliceCopy are not equal")
	}

	twoDSlice := make([][]int, 3) // create a 2D slice with 3 rows (outer slice with 3 elements, each being a slice)
	for i := 0; i < 3; i++ {      // iterate over the rows (outer slice indices)
		innerLen := i + 1                    // determine the length of the inner slice for the current row (increasing length)
		twoDSlice[i] = make([]int, innerLen) // create an inner slice with the calculated length and assign it to the current row
		for j := 0; j < innerLen; j++ {      // iterate over the elements of the inner slice
			twoDSlice[i][j] = i + j // assign a value to each element (sum of row index and column index)
			fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice at index %d\n", i+j, i, j) // print the value being added
		}
	}
	fmt.Println("2D Slice:", twoDSlice) // print the entire 2D slice to the console // Output: 2D Slice: [[0] [1 2] [2 3 4]]

	/*
		Key concepts:
	2D Slice Creation: make([][]int, 3) creates a slice of slices (a 2D slice) with 3 rows. Initially, each row is nil.
	Dynamic Inner Slices: Each row (inner slice) is created with a length that increases with the row index (i + 1).
	Nested Loops: The outer loop iterates over rows, and the inner loop iterates over the elements of each row.
	Value Assignment: Each element in the 2D slice is assigned a value based on its position (i + j).
	*/

	// slice[low:high] // slice of array that will take elements from index low to high-1

	slice2 := slice1[2:4]
	fmt.Println(slice2)

	fmt.Println("The capacity of slice2 is:", cap(slice2)) // Output: The capacity of slice1 is: 5
	fmt.Println("The length of slice2 is:", len(slice2)) // Output: The length of slice1 is: 2
}
