package main

import "fmt"

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func main() {
	fmt.Println("Welcome to defer Study!")
	//defer= delay the execution of a statement until function exits
	defer fmt.Println("Shykat") //defer keyword is used to delay the execution of a statement until function exits
	defer fmt.Println("Md")
	fmt.Print("Abir\n")
	myDefer()
}
