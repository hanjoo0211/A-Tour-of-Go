package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "세상"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	// Truth = false
}
