package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for {
		sum += 1
		fmt.Println(sum)
		if sum > 10000000 {
			fmt.Println("break")
			break
		}
	}
}
