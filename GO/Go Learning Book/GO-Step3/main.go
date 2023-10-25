package main

import "fmt"
func change(val int){
	val = 25
}
func callByRef(ptr *int){
	*ptr = 25
}

func main() {
	x := 10
	change(x)
	fmt.Println(x)
	callByRef(&x)
	fmt.Println(x)
	
	
	
	/*
	//pointer
	var p *int
	x:= 10
	p = &x
	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(*p)
	*p = *p -1
	fmt.Println(x)
*/
}




/*
func add(x,y float32)float32{
	return x+y
}
func sub(x,y float32)float32{
	return x-y
}
func multi(x,y float32)float32{
	return x*y
}
func div(x,y float32)float32{
	return x/y
}


func main() {
	var num1, num2, result float32
	var option string
	fmt.Println("Welcome to calculator")
	i := true
	for i == true {
		fmt.Println("Enter first number:")
		fmt.Scanln(&num1)
	
		fmt.Println("Enter second number:")
		fmt.Scanln(&num2)
	
		fmt.Println("Choose an option: +, -, *, /")
		fmt.Scanln(&option)
	
		switch option {
		case "+":
			result = add(num1, num2)
		case "-":
			result = sub(num1, num2)
		case "*":
			result = multi(num1, num2)
		case "/":
			result = div(num1, num2)
		default:
			fmt.Println("Invalid option")
			continue
		}
		fmt.Println("Result is:", result)
	}
}
*/

/*
// function to display message
func displayMsg(countryName string) {
	fmt.Println("I am from:", countryName)
}
func squre(number int){
	result := number * number
	fmt.Println("result is:", result)
}

func add (number int, number1 int){
	result := number + number1
	fmt.Println("result is:", result)	
}

func sub (number int, number1 int) int{
	result := number - number1
	return result	
}
func msg()string{
	return "Hello Dog"
}

func main() {
	displayMsg("Bd")
	displayMsg("India")
	displayMsg("USA")
	squre(5)
	add(5, 10)
	fmt.Printf("result is: %v\n", sub (10, 5))
	text := msg()
	fmt.Println(text)
}
*/