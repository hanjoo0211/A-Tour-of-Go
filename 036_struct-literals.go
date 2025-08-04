package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
	p.X = 100
	fmt.Println(v1, p, v2, v3)
	v1.X = 150
	fmt.Println(v1, p, v2, v3)
	p2 := &v2
	p2.X = 200
	fmt.Println(v1, p, v2, v3)
}
