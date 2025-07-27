package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("10")
	defer fmt.Println("11")

	fmt.Println("done")
}
