package main

import "fmt"

func editMap(mp map[string]int) {
	mp["ONE"] = 10
}

func main() {
	var mp map[string]int
	mp["ONE"] = 1

	editMap(mp)
	fmt.Print(mp)
}
