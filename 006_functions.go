package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addStr(x string, y string) string {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(addStr("wow", "hi"))
}
