package main

import "fmt"

func add(x int, y int)int { //int is return type
	return x+y
}

func greet(){
	fmt.Println("Salam Alaikum")
	
}
func proAdd(values ...int)int{ //variadic function (multiple arguments) //values is slice of int
	total := 0                  //... is used to pass multiple arguments
	for _,val := range values{
		total += val
	}
	return total
}
func proSub(values ...int)(int,string){
	total :=0
	for _,val := range values{
		total -= val
	}
	return total,"Subtraction is done"
}

func main() {
	fmt.Println("Welcome to function Study!")
	greet()
	var num1,num2 int
	fmt.Println("Enter Two Numbers")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	result := add(num1,num2)
	fmt.Println("Result is ",result)
	proRes := proAdd(20,30,50)
	fmt.Println("Result is ",proRes)

	proRes1,doneMsg := proSub(300,100,500)
	fmt.Println("Result is ",proRes1)
	fmt.Println(doneMsg) //doneMsg is second return value 

}