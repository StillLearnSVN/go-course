package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariable map[keyType]valueType
	// mapVariable = make(map[keyType]valueType)

	// using a map Literal
	// mapVariable := map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2,
	// }

	myMap := make(map[string]int)

	fmt.Println(myMap)

	myMap["key1"] = 1
	myMap["code"] = 2

	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])

	// To check if a key exists in the map
	_, ok := myMap["key1"]
	if ok {
		fmt.Println("Key exists with key1")
	} else {
		fmt.Println("Key does not exist with key1")
	}

	myMap2 := map[string]int{"a": 1, "b": 2, "c": 3}
	myMap3 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(myMap2)

	if maps.Equal(myMap2, myMap3) {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

	for _, v := range myMap3 {
		fmt.Println("Value:", v)
	}

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value")
	} else {
		fmt.Println("The map is not initialized to nil value")
	}

	val := myMap4["key"]
	fmt.Println(val)

	// myMap4["key"] = "value"
	// fmt.Println(myMap4)

	myMap4 = make(map[string]string)
	myMap4["key"] = "value"
	fmt.Println(myMap4)

	fmt.Println("myMap length is", len(myMap))
	fmt.Println("myMap4 length is", len(myMap4))

	myMap5 := make(map[string]map[string]string) //  a map where the keys are strings, and the values are maps with string keys and string values.

	myMap5["map1"] = myMap4
	fmt.Println(myMap5)

	/*
	~The code then demonstrates the creation of a more complex map, myMap5, using the make function. This map has string keys, 
	but its values are themselves maps with string keys and string values (map[string]map[string]string). 
	This structure allows for nested mappings, enabling hierarchical or grouped data storage.

		The myMap4 map is assigned as a value to the key "map1" in myMap5. This effectively nests myMap4 within myMap5. 
		Finally, the contents of myMap5 are printed, showing a map where the key "map1" maps to the entire myMap4 map. 
		This demonstrates how maps can be used to build more complex data structures in Go.
	*/
}
