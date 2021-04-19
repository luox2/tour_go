package main

import "fmt"

type vertex struct {
	X, Y float64
}

func (v vertex) Abs() float64 {
	return v.X*v.X + v.Y*v.Y
}

func main() {
	fmt.Println(vertex{3, 4}.Abs())
}
