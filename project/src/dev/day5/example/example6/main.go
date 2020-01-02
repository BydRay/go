package main

import (
	"fmt"
)

type Car1 struct {
	name string
	age  int
}

type Car2 struct {
	name string
	age  int
}

type Train struct {
	Car1
	Car2
}

func main() {
	var t Train

	t.Car1.name = "train"
	t.Car2.age = 100

	fmt.Println(t)

}
