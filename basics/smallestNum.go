package main

import "fmt"

func main() {
	list := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, -97, 9, 17}
	smlNum := list[0]
	for i := range list {
		if list[i] < smlNum {
			smlNum = list[i]
		}
	}
	fmt.Println(smlNum)

}
