package main

import "fmt"

func main() {
	fmt.Println("Welcome to ifelse Study!")

	loginCount := 9

	var result string

	if loginCount < 10{
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else{
		result = "Exactly 10 login count"
	}
	fmt.Println("Result is", result)
	if 9%2 ==0{ 
		fmt.Println("Number is even")
	} else{
		fmt.Println("Number is odd")
	}
	if num := 3; num <10{ //num is only available inside this if block
		fmt.Println("Num is less than 10")
	} else{
		fmt.Println("Num is not less than 10")
	}
}