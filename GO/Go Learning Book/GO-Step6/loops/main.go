package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops Study!")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	// for i := 0; i<len(days); i++{ 
	// 	fmt.Println(days[i])
	// }

	// for i := range days{ //range returns index of the element of the array or slice 
	// 	fmt.Println(days[i])
	// }
	// for index, day := range days{
	// 	fmt.Printf("Index: %v, Day: %v\n", index, day)
	// }
	rougueValue := 1
	for rougueValue < 10{ //while loop
		if rougueValue == 2{
			goto lco //goto lco label when rougueValue is 2 
		}
		if rougueValue == 5{
			break //break the loop when rougueValue is 5 
		}
		fmt.Println(rougueValue)
		rougueValue++
	}
	fmt.Print("Done\n\n")

	lco: //lco label //goto lco label when rougueValue is 2 //lco is a label name
	fmt.Println("Welcome to lco") 
}