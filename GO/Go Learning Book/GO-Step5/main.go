package main

import "fmt"

func main() {
	
	//dynamic array
	var student  [] string //slice
	var studentName string
	var numOfStudent int
	fmt.Println("Enter the number of student: ")
	fmt.Scanln(&numOfStudent)

	for i := 0; i <numOfStudent; i++ {
		fmt.Println("Enter the name of student: ")
		fmt.Scanln(&studentName) 
		student = append(student, studentName) //append function
	}
	fmt.Println(student)
	fmt.Println(len(student)) //length of array
	fmt.Println(cap(student)) //capacity of array
	
	
	
	
	/*
	var dest = [6]string{"Finland", "India", "USA", "UK", "BD", "China"} //array

	fmt.Println(dest[3]) //accessing array element
	fmt.Print("\n")

	for i := 0; i < len(dest); i++ { //printing array element
		fmt.Println(dest[i])
	}
	fmt.Print("\n")
	fmt.Println("Enter Your choise: 0 to 4")
	var choise int
	fmt.Scanln(&choise)
	fmt.Println(dest[choise]) 

	var name [5]string
	fmt.Println("Enter 5 name: ")
	for i := 0; i < len(name); i++ { //taking input
		fmt.Scanln(&name[i])
	}
	fmt.Printf("1st name: %v, 2nd name: %v, 3rd name: %v, 4th name: %v, 5th name: %v\n", name[0], name[1], name[2], name[3], name[4])
	*/
}
