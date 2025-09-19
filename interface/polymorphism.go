package main

import "fmt"

type Shape interface {
	Area()
}

type Rectangle struct {
	length int
	width  int
}
type Circle struct {
	radius int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() int {
	return c.radius * c.radius
}

func main() {
	r := Rectangle{
		length: 7,
		width:  7,
	}
	fmt.Println(r.Area())
	c := Circle{
		radius: 20,
	}
	fmt.Println(c.Area())
}
