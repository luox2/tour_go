package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["bell lab"] = Vertex{40.68433, -74.39967}
	fmt.Println(m, m["bell lab"])
	fmt.Println(m["hello world"])
	e, ok := m["hell world"]
	fmt.Println(e, ok)

	n := map[string]Vertex{
		"hello": {1, 2},
		"world": {3, 4},
	}
	fmt.Println(n)
}
