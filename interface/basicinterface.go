package main

import "fmt"

type Shape interface {
	Area() int
}

type Rect struct {
	length  int
	breadth int
}

func (r Rect) Area() int {
	return r.length * r.breadth
}
func nothing(s Shape) {
	fmt.Println("good morning")
}

func main() {
	r := Rect{12, 13}
	fmt.Println(r.Area())
	nothing(r)
}
