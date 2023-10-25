package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Printf("Type of fruits is %T\n", fruits)

	fruits = append (fruits,   "papaya")
	fmt.Println(fruits)

	fruits = append(fruits[1:3])  //remove the first and last element of fruits slice and store the result in fruits slice itself 
	fmt.Println(fruits)

	highScr := make([]int,4) //create a slice with 4 elements and store in highScr variable //make([]int,4) is a shorthand for []int{0,0,0,0}
	highScr[0] = 234
	highScr[1] = 945
	highScr[2] = 284
	highScr[3] = 904

	highScr = append (highScr, 555,666,777) //add 555,666,777 to highScr slice and store the result in highScr slice itself //append() function can be used to add one or more elements to a slice

	fmt.Println(highScr)

	fmt.Println(sort.IntsAreSorted(highScr)) //check if highScr slice is sorted in ascending order
	sort.Ints(highScr) //sort highScr slice in ascending order
	
	fmt.Println(highScr) 
		fmt.Print("Done\n\n")

	// how to remove an element from a slice based on index in Go

	var courses = [] string{"Go", "Java", "Swift", "C","C++" }
	fmt.Println(courses)
	var index int = 2 //index of the element to be removed
	courses = append(courses[:index], courses[index+1:]...) //remove the element at index 2 from courses slice and store the result in courses slice itself
		                                                    //append() function can be used to add one or more elements to a slice
															//append() function can also be used to remove one or more elements from a slice
															//... is a variadic parameter which means it can take zero or more arguments
											
	fmt.Println(courses)
}