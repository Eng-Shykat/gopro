// A go program this is a comment
package main

import "fmt"

/*
This is a multi-line comment
Date: 2023-06-09
*/

func main() {
	/*fmt.Println("Hello, World!")
	fmt.Println("This is my first Go program.")
	fmt.Println("I am Abir")
	fmt.Println("I am 26 years old")
	fmt.Println("I am from Bangladesh")
	fmt.Println("Name \t\t Age")
	fmt.Println("Abir \t\t 27")
	fmt.Println("\"Shykat\"")
	*/
	//variable declaration
	var name string
	var age int
	var result float32
	var country string

	//variable initialization
	name = "Abir"
	age = 27
	result = 50.5
	country = "Bangladesh"

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Result: ", result)
	fmt.Println("Country: ", country)
	fmt.Println(name, "has got", result, "in the exam")
	fmt.Println(name, "is from", country)	

	//variable declaration and initialization
	var name1  = "Shykat"
	var age1  = 26
	var result1  = 90.5
	var country1  = "Bangladesh"
	fmt.Println("Name: ", name1)
	fmt.Println("Age: ", age1)
	fmt.Println("Result: ", result1)
	fmt.Println("Country: ", country1)

	//variable declaration and initialization 2
	name2  := "Abir Shykat"
	age2  := 27
	result2  := 50.54
	country2  := "Bangladesh"
	fmt.Println("Name: ", name2, "Age: ", age2, "Result: ", result2, "Country: ", country2)

/*fullName := "Abir Shykat"
age := 27
result := 50.54
const COUNTRY = "Bangladesh"
fmt.Printf("%v is a student\n", fullName)
fmt.Printf("%v is %v years old\n", fullName, age)
fmt.Printf("%v has got %v /100 in the exam\n", fullName, result)
fmt.Printf("%v is from %v\n", fullName, COUNTRY)
*/
var name12 string
var num1, num2 int
//var age int
//var result float32

fmt.Println("Enter your name: ")
fmt.Scan(&name12)
fmt.Println("Hello", name12)
fmt.Println("Enter two numbers: ")
fmt.Scan(&num1, &num2)
//fmt.Printf("num1 = %v, num2 = %v", num1, num2)

//fmt.Println("Enter 2 numbers: ")
//fmt.Scanf("%v %v", &num1, &num2)
fmt.Println("num1 = ", num1,"num2 = ", num2)
//Input: Scan, Scanf, Scanln
//Output: Print, Println, Printf

}

//uint8 0 to 255 (positive number)
//int8 -128 to 127 (negative number)
//float32 32 bit floating point number
//float64 64 bit floating point number(big number)
//complex64 complex number
//variable naming convention
//1:letters,numbers,underscore
//2:cannot start with a number
//3:case sensitive
//4:camelCase
//5:cannot use reserved keywords
//6:cannot use white space
//7:cannot use operators
//8:cannot use double quotes
//9:cannot use single quotes
//10:cannot use backticks
//11:cannot use backslash
//12:cannot use colon
//13:cannot use semicolon
//14:cannot use comma
