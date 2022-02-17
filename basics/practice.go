package main

import (
	"fmt"
)

func main() {
	var x interface{}
	x = "a"
	fmt.Println(x)
	x = 10
	fmt.Println(x)
}
