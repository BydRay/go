package main

import "fmt"

func main() {
	var a []int
	a = append(a, 2)
	a = append(a, a...)

	var b int
	fmt.Printf("%T\n", b)

	var arr [5]int
	arr2 := arr
	arr2[2] = 100 // arr1 不会改变

	fmt.Println(arr)
	fmt.Println(arr2)

	str := "asdf"
	// str[2] = 'a'
	fmt.Println(str)

	m1 := make(map[string]int, 0)
	m1 = make(map[string]int, 1)

	fmt.Println(m1)
	fmt.Println(len(m1))

}
