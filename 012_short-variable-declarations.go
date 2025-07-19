package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)

	a1, a2, a3 := "Hello", 3, true
	fmt.Println(a1, a2, a3)
}
