package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	v := Point{3, 4}
	fmt.Println(v)
	v.X = 5
	fmt.Println(v)
	p := &v
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(p.X)
	fmt.Println((*p).X)

}
