package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (V Vertex) Abs() float64 {
	return V.X*V.X + V.Y*V.Y
}

// methods pointer

func (V *Vertex) Scale(f float64) {
	V.X = V.X * f
	V.Y = V.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
}
