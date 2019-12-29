package main

func main() {
	var a []int
	a = append(a, 2)
	a = append(a, a...)
}
