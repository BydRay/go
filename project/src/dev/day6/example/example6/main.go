package main

import (
	"fmt"
)

type Student struct {
	Name string
	Sex  string
}

func Test(a interface{}) {
	b, ok := a.(Student)
	if ok == false {
		fmt.Println("convert faild")
		return
	}

	fmt.Println(b)
}

func just(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("%d params is bool, value is %v\n", index, v)
		case int, int32, int64:
			fmt.Printf("%d params is int, value is %v\n", index, v)
		case float32, float64:
			fmt.Printf("%d params is float, value is %v\n", index, v)
		case string:
			fmt.Printf("%d params is string, value is %v\n", index, v)
		case Student:
			fmt.Printf("%d params is student, value is %v\n", index, v)
		case *Student:
			fmt.Printf("%d params is *studhent, value is %v\n", index, v)
		}
	}
}

func main() {
	var b Student = Student{
		Name: "stu01",
		Sex:  "female",
	}
	Test(b)
	just(28, 8.2, "this is a test", b, &b)
}
