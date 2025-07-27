package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	a := 3
	for ; a < 10; a++ {
		if a%2 == 0 {
			fmt.Println(a, "is even")
		} else {
			fmt.Println(a, "is odd")
		}
	}
}
