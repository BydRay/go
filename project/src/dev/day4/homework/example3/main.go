package main

import (
	"fmt"
)

// func main() {

// 	str := "a2啊是"
// 	fmt.Println("len ", len(str))
// 	for i := 0; i < len(str); i++ {
// 		fmt.Printf("%d - %x\n", i, str[i])
// 	}

// 	t := []rune(str)
// 	fmt.Println(len(t))
// 	for i := 0; i < len(t); i++ {
// 		fmt.Printf("%d - %c - %x\n", i, t[i], t[i])
// 	}

// }

func process(str string) bool {
	t := []rune(str)
	length := len(t)
	for i, _ := range t {
		if i == length/2 {
			break
		}

		last := length - i - 1
		if t[i] != t[last] {
			return false
		}
	}

	return true
}

func main() {
	var str string
	fmt.Scanf("%d", &str)

	if process(str) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

}
