package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Println("Today is " + today.String())
	fmt.Println("Tomorrow is " + (today + 1).String())

	switch time.Wednesday {
	case today + 0:
		fmt.Println("Today is Wednesday.")
	case today + 1:
		fmt.Println("Tomorrow is Wednesday.")
	case today + 2:
		fmt.Println("In two days is Wednesday.")
	case today + 3:
		fmt.Println("In three days is Wednesday.")
	case today + 4:
		fmt.Println("In four days is Wednesday.")
	case today + 5:
		fmt.Println("In five days is Wednesday.")
	default:
		fmt.Println("Next Wednesday is too far away.")
	}
}
