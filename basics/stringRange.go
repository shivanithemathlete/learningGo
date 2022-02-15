package main

import "fmt"

func main() {
	s := "string"
	for i, c := range s {
		fmt.Printf("%v  %c %v\n", i, c, c)
	}
	s1 := "Go"
	s2 := " is awesome"
	res := fmt.Sprintf("%s %s", s1, s2)
	fmt.Printf(res)
}
