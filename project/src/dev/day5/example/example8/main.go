package main

import (
	"fmt"
)

type Car struct {
	weight int
	name   string
}

func (p *Car) Run() {
	fmt.Print("running")
}

type Bike struct {
	Car
	wheel int
}

type Train struct {
	c Car
}

func main() {
	var a Bike
	a.weight = 100
	a.name = "bike"
	a.wheel = 2

	fmt.Println(a)
	a.Run()

	var b Train
	b.c.weight = 100
	b.c.name = "train"
	b.c.Run()

}
