package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

func (r Rectangle) Area() int {
	return r.length * r.breadth
}

func main() {
	r := Rectangle{
		length:  23,
		breadth: 30,
	}
	fmt.Println(r.Area())

}
