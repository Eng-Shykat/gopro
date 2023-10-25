package main

import (
	"fmt"
	"bufio" 
	"os"
)
const Author string = "Shykat" //public constant variable

func main(){
	//bufio > buffer input/output
	welcome := "Welcome to Go!"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin) //read from standard input (keyboard) and store in reader variable 
	fmt.Println("Enter your full name: ")
	input, _ := reader.ReadString('\n') //read string until new line character and store in input variable 
	fmt.Println("Your name is: ", input)
}