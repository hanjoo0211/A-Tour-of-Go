package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func swapTriple(x, y string, z int) (string, int, string) {
	return y, z, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	c, d, e := swapTriple("hi", "wow", 13)
	fmt.Println(c, d, e)
}
