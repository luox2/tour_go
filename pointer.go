package main

import "fmt"

func main() {
	i, j := 27, 2701

	p := &i
	var q *int = &i
	fmt.Println(*p)
	fmt.Println(*q)

	q = &j
	*q = *q / 100
	fmt.Println(*q, q)
}
