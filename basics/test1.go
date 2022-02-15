package main

import "fmt"

func Reduce(x, y int, f func(a, b int) int) int {
	var res int
	res = f(x, y)
	return res
}
func Addition(x, y int) int {
	return x + y
}
func Subtraction(x, y int) int {
	return x - y
}
func Multiplication(x, y int) int {
	return x * y
}
func Division(x, y int) int {
	return x / y
}
func main() {
	x, y := 20, 10
	fmt.Println("Additon", Reduce(x, y, Addition))
	fmt.Println("Subtraction", Reduce(x, y, Subtraction))
	fmt.Println("Multiplication", Reduce(x, y, Multiplication))
	fmt.Println("Division", Reduce(x, y, Division))
}
