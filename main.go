package main

import (
	"fmt"

	a "github.com/shivanithemathlete/learningGo/pkg"
)

func Square(x int) int {
	return x * x
}
func main() {
	fmt.Println("Calling Reduce from package pkg")
	inputSlice := []int{2, 4, 3, 9, 0, 1, 4}
	fmt.Println(a.Reduce(inputSlice, Square))

}
