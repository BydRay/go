package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3}
	// fmt.Printf("%T\n", a)

	var b []*int
	for _, v := range a {
		// fmt.Println(i, &v)
		b = append(b, &v)
	}

	fmt.Println(b)
	fmt.Println(*b[0])
}
