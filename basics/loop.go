package main

import "fmt"

func main() {
	sum := 1
	for sum {
		sum += sum
	}
	fmt.Println(sum)
}
