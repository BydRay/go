package main

import "fmt"

type integer int

type Student struct {
	Number int
	Score  float32
}

type Stu Student

func main() {
	var i integer = 1000
	var j int = 100

	j = int(i)
	fmt.Println(j)

	var a Student
	a = Student{
		Number: 30,
		Score:  32,
	}

	var b Stu
	a = Student(b)
	fmt.Println(a)
}
