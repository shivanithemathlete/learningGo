package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name of Person: %s \nAge of Person: %v", p.Name, p.Age)
}

func main() {
	a := Person{"Tont Stark", 50}
	fmt.Println(a)
}
