package main

import "fmt"

func main() {

	//break and continue statement

	for i:= 1; i<=10; i++{
		if i % 2 == 0{
		continue
	}
	fmt.Printf("%v\n",i)
	}
		fmt.Printf("Done\n\n\n")

	for i:= 1; i<=10; i++{
		if i % 2 == 0{
		break
	}
	fmt.Printf("%v\n",i)
	}
	
	
	
	
	/*
	//series 1,2,3,4,5,6,7,8,9,10
	var endNum, startNum int
	fmt.Printf("Enter the first number of Series: ")
	fmt.Scan(&startNum)
	fmt.Printf("Enter the last number of Series: ")
	fmt.Scan(&endNum)

	sum := 0
	for i := startNum; i <= endNum; i++ {
		sum += i // sum = sum + i
	}
	fmt.Println(sum)

	//series 2+4+6+8+10...n

	var endNum1, startNum1 int
	fmt.Println("Enter The 1st Number")
	fmt.Scan(&startNum1)
	fmt.Println("Enter The 2nd Number")
	fmt.Scan(&endNum1)

	sum1 := 0
	for i := startNum1; i <= endNum1; i += 2 {
		sum1 += i
	}
	fmt.Println(sum1)

	*/

	/*
		for x := 2; x<=10; x+=2{
			fmt.Printf("%v\n",x)
		}
		fmt.Printf("Done\n\n")
		for x := 1; x<=100; x++{
			if x % 2 == 0{
				fmt.Printf("Even %v\n",x)
			}else {
				fmt.Printf("Odd %v\n",x)
			}
		}
	*/

	/*
		//op1
		//for loop > init; condition; post
		for x := 1; x <=10; x++{
			fmt.Printf("%v\n",x)
		}
		fmt.Printf("Done\n")

		//op2
		y := 10
		for y >= 1{
			fmt.Printf("%v\n",y)
			y--
		}
	*/

	/*
		var classNum int
		fmt.Printf("Enter your class number: ")
		fmt.Scan(&classNum)

		switch classNum {
		case 1,2,3,4,5:
			fmt.Printf("Primary School")
		case 6,7,8,9,10:
			fmt.Printf("High School")
		case 11,12:
			fmt.Printf("College")
		default:
			fmt.Printf("Invalid input")
		}
	*/

	/*
		//disit spelling using switch
		var num int
		fmt.Printf("Enter a number: ")
		fmt.Scan(&num)

		switch num {
		case 0:
			fmt.Printf("Zero")
		case 1:
			fmt.Printf("One")
		case 2:
			fmt.Printf("Two")
		case 3:
			fmt.Printf("Three")
		case 4:
			fmt.Printf("Four")
		case 5:
			fmt.Printf("Five")
		case 6:
			fmt.Printf("Six")
		case 7:
			fmt.Printf("Seven")
		case 8:
			fmt.Printf("Eight")
		case 9:
			fmt.Printf("Nine")
		default:
			fmt.Printf("Invalid input")
		}
	*/

	/*
		//disit spelling
		var num int
		fmt.Printf("Enter a number: ")
		fmt.Scan(&num)

		if (num == 0){
			fmt.Printf("Zero")
		}else if (num == 1){
			fmt.Printf("One")
		}else if (num == 2){
			fmt.Printf("Two")
		}else if (num == 3){
			fmt.Printf("Three")
		}else if (num == 4){
			fmt.Printf("Four")
		}else if (num == 5){
			fmt.Printf("Five")
		}else if (num == 6){
			fmt.Printf("Six")
		}else if (num == 7){
			fmt.Printf("Seven")
		}else if (num == 8){
			fmt.Printf("Eight")
		}else if (num == 9){
			fmt.Printf("Nine")
		}else {
			fmt.Printf("Invalid input")
		}
	*/

	/*
		var num1,num2,num3 int
		fmt.Printf("Enter three numbers: ")
		fmt.Scan(&num1,&num2,&num3)

		if num1 > num2 && num1 > num3{
			fmt.Printf("Leargest number is %v",num1)
		}else if num2 > num1 && num2 > num3{
			fmt.Printf("Leargest number is %v",num2)
		}else if num3 > num1 && num3 > num2{
			fmt.Printf("Leargest number is %v",num3)
		}else {
			fmt.Printf("Numbers are equal")
		}
	*/

	/*
		var num1,num2,num3 int
		fmt.Printf("Enter three numbers: ")
		fmt.Scan(&num1,&num2,&num3)

		if num1 > num2{
			if num1 > num3{
				fmt.Printf("Leargest number is %v",num1)
			}else {
				fmt.Printf("Leargest number is %v",num3)
			}
		}else if num2 > num1{
			if num2 > num3{
				fmt.Printf("Leargest number is %v",num2)
			}else {
				fmt.Printf("Leargest number is %v",num3)
			}
			}else {
				fmt.Printf("Numbers are equal")
			}
	*/

	/*
		var num1,num2 int
		fmt.Printf("Enter two numbers: ")
		fmt.Scan(&num1,&num2)

		if num1 > num2 {
			fmt.Printf("%v is greater than %v",num1,num2)
		} else if num1 < num2 {
			fmt.Printf("%v is greater than %v",num2,num1)
		} else if num1 == num2 {
			fmt.Printf("%v is equal to %v",num1,num2)
		} else {
			fmt.Printf("Invalid input")
		}

	*/

	/*
		//even and odd

		var num int

		fmt.Printf("Enter a number: ")
		fmt.Scan(&num)

		if num % 2 == 0 {
			fmt.Printf(" Your entered number is even")
		} else {
			fmt.Printf("Your entered number is odd")
		}

	*/

	/*
		// control statements - if, else, else if, switch

		n := 0

		if n > 0{
			fmt.Printf("Positive")
		}else if n < 0{
			fmt.Printf("Negative")
		}else {
		fmt.Printf("Zero")
		}
	*/
}
