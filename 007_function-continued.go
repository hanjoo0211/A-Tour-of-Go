package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func addPrint(x, y int, z string) bool {
	fmt.Println(x + y, z)
	return true
}

func main() {
	fmt.Println(add(42, 13))
	// fmt.Println(add("hi ", "hello"))
	addPrint(50, 20, "wow")
}
