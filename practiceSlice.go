package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 5)
	fmt.Println(a, b)
	b = a[:3]
	fmt.Println(b)
	b[0] = 6
	fmt.Println(a, b)
}
