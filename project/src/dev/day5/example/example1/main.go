package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	var stu Student

	stu.Name = "hua"
	stu.Age = 18
	stu.Score = 80

	var stu1 *Student = &Student{
		Age:  20,
		Name: "hua",
	}

	var stu3 = Student{
		Age:  20,
		Name: "hua",
	}
	fmt.Println(stu1.Name)
	fmt.Println(stu3)
	fmt.Printf("Name:%p\n", &stu.Name)
	fmt.Printf("Age:%p\n", &stu.Age)
	fmt.Printf("score:%p\n", &stu.Score)

}
