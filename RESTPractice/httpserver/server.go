package main

import (
	"fmt"
	"log"
	"net/http"
) 
func main(){

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello Server")	
	})

	const port string = ":8080"
	
	fmt.Println("Server Listening on port:", port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatalln("error starting server", err)
	}
}
