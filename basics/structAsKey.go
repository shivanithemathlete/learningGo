package main

import "fmt"

type Address struct {
	Name    string
	City    string
	Pincode int
}

func main() {
	a1 := Address{"abc", "xyz", 473838}
	a2 := Address{"pqr", "rst", 573929}
	var mp map[Address]int
	if mp == nil {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	sample := map[Address]int{a1: 1, a2: 2}
	sample[Address{"edueh", "eruwyej", 378290}] = 8
	fmt.Println(sample)
}
