package main

import "fmt"

func main() {
	sum := 1
	for sum < 10000000000000 {
		sum += sum
	}
	fmt.Println(sum)
}
