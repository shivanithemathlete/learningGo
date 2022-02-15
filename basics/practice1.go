package main

import "fmt"

func main() {
	inputString := "MISSISSIPPI"
	myMap := make(map[string]int)
	for _, c := range inputString {
		myMap[string(c)]++
	}
	for key, val := range myMap {
		fmt.Printf("Key: %v, Val: %v\n", key, val)
	}
}
