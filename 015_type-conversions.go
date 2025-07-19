package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	a := 3
	fmt.Printf("Type: %T Value: %v\n", a, a)
	var b float64 = float64(a)
	fmt.Printf("Type: %T Value: %v\n", b, b)
	var c uint = uint(b)
	fmt.Printf("Type: %T Value: %v\n", c, c)
}
