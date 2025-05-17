package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{Name: "Alice", Email: "alice@example.com"}
	fmt.Println(user)
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))

	var user1 User
	err = json.Unmarshal(jsonData, &user1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User created from json data", user1)

	jsonData1 := `{"name": "John", "email": "john@example.com"}`
	reader := strings.NewReader(jsonData1)
	decoder := json.NewDecoder(reader)

	var user2 User
	err = decoder.Decode(&user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user2)

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	err = encoder.Encode(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encoded json string:", buf.String())

}

// Serialize-Deserialize means converting a data structure into a format that can be easily stored or transmitted and then reconstructing it later.
// In Go, this is commonly done using the encoding/json package, which provides functions for encoding (marshaling) and decoding (unmarshaling) JSON data.
// The code demonstrates how to serialize a struct to JSON and deserialize JSON back into a struct using the json package in Go.
// The User struct is defined with JSON tags to specify how the fields should be named in the JSON output.

// Serialize apa dan Deserialize:
// Serialize adalah proses mengubah data menjadi format tertentu (seperti JSON) untuk penyimpanan atau transmisi.
// Deserialize adalah proses mengubah data dari format tersebut kembali menjadi struktur data asli.
// Serialize-Deserialize adalah proses mengubah struktur data menjadi format yang dapat dengan mudah disimpan atau ditransmisikan,
// dan kemudian merekonstruksinya kembali nanti. Dalam konteks Go, ini biasanya dilakukan menggunakan paket encoding/json,
// yang menyediakan fungsi untuk encoding (marshalling) dan decoding (unmarshalling) data JSON.

/*
This Go program demonstrates how to work with JSON data using Go’s standard library. It defines a simple User struct with Name and Email fields,
each tagged for JSON serialization. In the main function, it first creates a User instance and prints it using fmt.Println, which outputs the
struct in Go’s default format.
Next, it serializes the User struct to JSON using json.Marshal, which converts the struct into a JSON-formatted byte slice. If marshaling fails,
the program logs the error and exits. The resulting JSON string is printed.

The program then demonstrates deserialization by unmarshaling the JSON data back into a new User struct using json.Unmarshal.
Again, error handling ensures the program exits if something goes wrong. The newly created struct is printed to confirm the data
was correctly parsed.

To show another way to decode JSON, the code uses a json.Decoder with a strings.NewReader to parse a JSON string into a User struct.
This approach is useful for reading JSON from streams or files.

Finally, the program demonstrates encoding a struct to JSON using a json.Encoder writing to a bytes.Buffer. This method is helpful when
you want to write JSON to an output stream or buffer. The encoded JSON string is printed, showing the result of the encoding process.

Overall, the code provides a clear example of serializing and deserializing Go structs to and from JSON, using both the Marshal/Unmarshal
functions and the Encoder/Decoder types for different use cases.
*/
