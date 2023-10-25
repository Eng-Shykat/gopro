package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch Study!")
	//rand.Seed(time.Now().UnixNano()) //generate a random number based on current time
	rand.NewSource(time.Now().Unix()) //generate a random number based on current time and store in rand variable
	diceNum := rand.Intn(6) +1 //generate a random number between 1 and 6 and store in diceNum variable
	fmt.Println("Value of diceNum is", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spots")
		fallthrough //fallthrough to next case
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough //fallthrough to next case
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move to 6 spots and roll again")
	default:
		fmt.Println("What was that!")
	}

}