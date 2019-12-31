package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

func insertHead(p *Student) {
	//var tail = p
	// for i := 0; i < 10; i++ {
	stu := Student{
		Name:  fmt.Sprintf("stu%d", 0),
		Age:   rand.Intn(100),
		Score: rand.Float32() * 100,
	}

	stu.next = p
	fmt.Printf("p = %p   p.next = %p\n", p, p.next)
	*p = stu
	fmt.Printf("(*p).Name = %s  (*p).next =  %p  (*p).next.next = %p\n", (*p).Name, (*p).next, (*p).next.next)

	return

	// }
}

// func trans(p *Student) {
// 	for p != nil {
// 		fmt.Println(*p)
// 		p = p.next
// 	}

// 	fmt.Println()
// }

func main() {
	var head *Student = new(Student)

	head.Name = "hua"
	head.Age = 18
	head.Score = 100

	insertHead(head)
	// trans(head)
}
