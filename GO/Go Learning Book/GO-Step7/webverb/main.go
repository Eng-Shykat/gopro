package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to webverb")
}

func PerformGetRequest() {
	const url = "http://localhost:3000/get"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
}