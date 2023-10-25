package main

import "fmt"

type User struct { //create a struct type and store in User variable
	Name string
	Email string
	Age int
	Status bool
}

func main() {
	fmt.Println("Welcome to struct Study!")
	//no inheritance in Go; No super or base keyword in Go; No extends keyword in Go

	shykat := User{"Shykat", "shykat@gmail.com", 26, true}
	fmt.Println(shykat)
	fmt.Printf("Details of shykat is %+v\n", shykat)

}