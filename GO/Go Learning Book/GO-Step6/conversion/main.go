package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Welcome to our service!")
	fmt.Println("Please rate our service (1-5): ")

	reader := bufio.NewReader(os.Stdin) //read from standard input (keyboard) and store in reader variable
	input, _ := reader.ReadString('\n') //read string until new line character and store in input variable

	fmt.Println("Thank you for rating our service with: ", input) //input is a string type variable 

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //convert string to float64 type	and store in numRating variable

	if err != nil{ //if error is not nil (not empty) then print error
		fmt.Println(err) //print error
	} else {
		fmt.Println("Added 1 to your rating: ", numRating + 1) //if no error then add 1 to numRating variable and print
}
}