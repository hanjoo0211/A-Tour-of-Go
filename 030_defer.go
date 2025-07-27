package main

import "fmt"

func main() {
	defer fmt.Println("world")

	sum := plus(1, 2)
	fmt.Println(sum)
	fmt.Println("hello")
}

func plus(a, b int) int {
	defer fmt.Println("plus called")
	fmt.Println("plus function is executing")
	return a + b
}
