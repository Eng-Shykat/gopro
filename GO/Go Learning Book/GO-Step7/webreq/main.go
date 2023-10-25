package main

import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	//"os"
)

const url = "https://innovatpark.com"

func main() {
	fmt.Println("Web Request Example")
	response, err := http.Get(url) // get the url and return the response and error if any occurs during the request
	checkError(err)

	fmt.Println("Response Status Code:", response)
	fmt.Printf("Response type: %T\n", response)
	defer response.Body.Close() // close the response body when done reading from it to avoid resource leaks\
	databytes, err := io.ReadAll(response.Body)
	checkError(err) 

	content := string(databytes) // convert the bytes to string for printing to the console or writing to a file
	fmt.Println(content) // print the content to the console for now but you can write to a file if you want
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
