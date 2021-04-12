package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)

	primes := []int{2, 3, 5, 7, 11}
	fmt.Println(primes)

	var s []int = primes[1:3]
	fmt.Println(s)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	var m []int
	m = append(m, 1)
	fmt.Println(m)
	m = append(m, 2, 3, 4, 5)
	fmt.Println(m)

	var pow = []int{1, 2, 4}
	for i, v := range pow {
		fmt.Println(i, v)
	}

}
