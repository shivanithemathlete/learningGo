package main

import (
	"fmt"
	"strconv"
)

type MyError struct {
	String string
	Value  int
}

func (e MyError) Error() string {
	return fmt.Sprintf("Cannot Convert %s to integer ", e.String)
}
func main() {
	s := "xyz"
	v, err := strconv.Atoi(s)
	myErr := &MyError{s, v}
	if err == nil {
		fmt.Printf("String %s converted to Int %v", s, v)
	} else {
		fmt.Println(myErr)
	}
}
