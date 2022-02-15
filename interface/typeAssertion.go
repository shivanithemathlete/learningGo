package main

import "fmt"

type I interface {
	Area() float64
}

const Pi float64 = 3.14

type Rectangle struct {
	length  int
	breadth int
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {

	return float64(r.length * r.breadth)
}
func (c *Circle) Area() float64 {
	return float64(c.radius * c.radius * Pi)
}

func main() {
	var i1, i2 I
	r := Rectangle{10, 20}
	i1 = r
	fmt.Println(i1.Area())

	i2 = &Circle{3}
	fmt.Println(i2.Area())

	// type switch
	switch v := i1.(type) {
	case Rectangle:
		fmt.Printf("Value: %v Type: %T\n", v, v)
	case *Circle:
		fmt.Println("Value in i2", *v)
	default:
		fmt.Println("Type Unknown")
	}

}
