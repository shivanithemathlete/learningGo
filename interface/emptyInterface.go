package main

import "fmt"

func main() {
	var i interface{}

	i = 2
	fmt.Printf("%v %T\n", i, i)

	i = "string"
	fmt.Printf("%v %T", i, i)

}
