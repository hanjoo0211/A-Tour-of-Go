package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Vertex3D struct {
	X int
	Y int
	Z int
}

type Vertexes struct {
	Vertex1 Vertex
	Vertex2 Vertex3D
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex3D{1, 2, 3})

	vs := Vertexes{
		Vertex1: Vertex{X: 1, Y: 2},
		Vertex2: Vertex3D{X: 3, Y: 4, Z: 5},
	}
	fmt.Println(vs)
	fmt.Println(vs.Vertex1.X, vs.Vertex1.Y)
	fmt.Println(vs.Vertex2.X, vs.Vertex2.Y, vs.Vertex2.Z)
}
