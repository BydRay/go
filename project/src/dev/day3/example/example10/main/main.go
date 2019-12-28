package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "111"
	// fmt.Scanf("%s", &str)

	number, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("convert failed, err:", err)
		return
	}

	fmt.Printf("%T\n", nil)

	fmt.Println(number)

}
