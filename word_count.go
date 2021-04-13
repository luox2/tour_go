package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	m_list := strings.Fields(s)
	for _, word := range m_list {
		_, v := m[word]
		if v {
			m[word] += 1
		} else {
			m[word] = 1
		}
	}
	return m
}

func main() {
	fmt.Println(WordCount("This eBook is for the use of anyone anywhere at no cost and with\nalmost no restrictions whatsoever.  You may copy it, give it away or\nre-use it under the terms of the Project Gutenberg License included\nwith this eBook or online at www.gutenberg.org"))
}
