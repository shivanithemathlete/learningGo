package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{6, 7, 8}
	a = append(a, b...)
	fmt.Println(a)
}
