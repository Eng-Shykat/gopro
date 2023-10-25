package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time Study!")

	presentTime := time.Now() //get current time and store in presentTime variable
	fmt.Println("The time is: ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //format time to string and print it out in the specified format

	fmt.Print("Done\n")

	createdDate := time.Date(2024, time.January, 23, 10, 23, 0, 0, time.UTC)                //create a time variable and store in createdDate variable
	fmt.Println("Created date is: ", createdDate.Format("02-January-2006 15:04:05 Monday")) //format time to string and print it out in the specified format

	currentTime := time.Now().Nanosecond() //get current time in nanosecond and store in currentTime variable
	fmt.Println("Current time in nanosecond is: ", currentTime)

}
