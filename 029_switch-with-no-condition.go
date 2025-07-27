package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	switch {
	case 1 < 2:
		fmt.Println("1 is less than 2")
	case 2 < 1:
		fmt.Println("2 is less than 1")
	default:
		fmt.Println("Neither condition is true")
	}
}
