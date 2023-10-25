package main

import "fmt"

func main() {
	x := 10 // 1010
	y := 20 // 10100
	and := x & y // 1010 & 10100 = 10100 = 20
	or := x | y
	xor := x ^ y
	fmt.Printf("X & Y = %v\n", and)
	fmt.Printf("X | Y = %v\n", or)
	fmt.Printf("X ^ Y = %v\n", xor)

	
	
	
	
	/*
	x := 10
	x = x + 2
	fmt.Println(x)

	x += 5 // x = x + 5
	fmt.Printf("x= %v\n", x)

	x++ // x = x + 1
	y := 2 // y = y + 1
	x += y // x = x + y
	fmt.Print("x= ", x, "\n")
	fmt.Printf("x= %v\n", x>=y)

	//z := x > 3 && x>21
	//z := x>3 || x>25
	z := ! (x>3)
	fmt.Printf("z= %v\n", z)
	*/

	/*
	const PI float32 = 3.1416
	var length, width float32
	var base, height, radius float32
	fmt.Println("Area of circle, rectangle, triangle")
	fmt.Printf("\n <<<<<>>>><<<>>>><<<>>\n")

	fmt.Println("Enter base, height for triangle")
	fmt.Printf("Enter base= ")
	fmt.Scan(&base)
	fmt.Printf("Enter height= ")
	fmt.Scan(&height)

	area := 0.5 * base * height

	fmt.Printf("Area of triangle= %.2f\n", area)

	fmt.Println("Enter radius for circle")
	fmt.Printf("Enter radius= ")
	fmt.Scan(&radius)

	circle := PI * radius * radius

	fmt.Printf("Area of circle= %.2f\n", circle)
	
	fmt.Println("Enter length, width for rectangle")
	fmt.Printf("Enter length= ")
	fmt.Scan(&length)
	fmt.Printf("Enter width= ")
	fmt.Scan(&width)

	rectangle := length * width

	fmt.Printf("Area of rectangle= %.2f\n", rectangle)
	*/

	/*
	var num1, num2 int
	fmt.Printf("num1= ")
	fmt.Scan(&num1)
	fmt.Printf("num2= ")
	fmt.Scan(&num2)

	result := num1 + num2
	fmt.Println("Result= ", result)
	fmt.Printf("%v + %v = %v\n", num1, num2, result)
	
	result = num1 - num2
	fmt.Printf("%v - %v = %v\n", num1, num2, result)

	result = num1 * num2
	fmt.Printf("%v * %v = %v\n", num1, num2, result)

	result = num1 / num2
	fmt.Printf("%v / %v = %v\n", num1, num2, result)

	result = num1 % num2
	fmt.Printf("%v %% %v = %v\n", num1, num2, result)

	var result2 float32 = float32(num1) / float32(num2)
	fmt.Printf("%v / %v = %.2f\n", num1, num2, result2)
	*/
	
	/*
	//string formatiing
	var name string = "Shykat"
	fmt.Printf("%s\n", name)
	fmt.Printf("%q\n", name)

	//float formatting
	var number = 3.1416
	fmt.Printf("%.2f\n", number)
	*/
	
	
	/*var decimalNumber int

	fmt.Printf("Enter a decimal number=  ")
	fmt.Scanf("%v", &decimalNumber)

	fmt.Printf("Binary Number %b", decimalNumber)
	fmt.Printf("\nOctal Number %o", decimalNumber)
	fmt.Printf("\nHexa Number %x", decimalNumber)
	*/

}
//type of operator
//Arithmetic operator>> + - * / % ++ --
//Relational operator>> == != > < >= <=
//Logical operator>> && || !
//Bitwise operator>> & | ^ << >>
//Assignment operator>> = += -= *= /= %= <<= >>= &= ^= |=
//Miscellaneous operator>> & * -> . sizeof() ?: (type) (type){} ,
//Unary operator>> ++ -- & * + - ! ~ sizeof() (type) (type){}
//Ternary operator>> ?:
//Binary operator>> + - * / % << >> & | ^ && || == != > < >= <=
//Special operator>> sizeof() & * -> . (type) (type){} ?: , (comma)
//area of circle = pi * r * r
//area of rectangle = length * width
//area of triangle = 1/2 * base * height
//area of square = side * side
//area of parallelogram = base * height
//area of trapezium = 1/2 * (a+b) * h
//area of rhombus = 1/2 * d1 * d2
//area of regular hexagon = 3/2 * a * a * sqrt(3)
//area of regular octagon = 2 * (1 + sqrt(2)) * a * a
//area of equilateral triangle = sqrt(3)/4 * a * a
//area of regular polygon = 1/2 * n * a * a * sin(360/n)
//area of ellipse = pi * a * b
//area of sector = 1/2 * r * r * theta
//area of segment = 1/2 * r * r * (theta - sin(theta))
//area of cyclic quadrilateral = sqrt((s-a) * (s-b) * (s-c) * (s-d))
//area of cyclic pentagon = sqrt((s-a) * (s-b) * (s-c) * (s-d) * (s-e))

