package main

import "fmt"

func returnSomething() (a string, b string) {
	return "a", "b"
}

func main() {
	fmt.Print(returnSomething())
}
