package main

import "fmt"

func main() {
	month := "january"
	switch month {
	case "january":
		fmt.Println("Winter.")
	case "december", "march":
		fmt.Println("Spring.")
	case "april", "may", "june":
		fmt.Println("Summer.")
	}
}
