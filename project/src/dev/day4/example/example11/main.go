package main

import (
	"fmt"
	"sort"
)

func testMapSort() {
	var a map[int]int
	a = make(map[int]int, 5)

	a[8] = 18
	a[3] = 13
	a[4] = 14
	a[5] = 15
	a[6] = 16

	var keys []int
	for k, _ := range a {
		keys = append(keys, k)
		// fmt.Println(k, v)
	}
	sort.Ints(keys)

	for _, v := range keys {
		fmt.Println(v, a[v])
	}

}

func testMapSort1() {
	var a map[string]int
	var b map[int]string

	a = make(map[string]int, 5)
	b = make(map[int]string, 5)

	a["abc"] = 101
	a["efg"] = 10

	for k, v := range a {
		b[v] = k
	}

	fmt.Println(b)

}

func main() {
	testMapSort()
	testMapSort1()
}
