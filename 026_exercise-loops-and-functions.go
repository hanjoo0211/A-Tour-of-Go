package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for originalZ := 0.0; math.Abs(z-originalZ) >= 1e-8; {
		originalZ = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z, originalZ, math.Abs(z-originalZ))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1000000))
}
