package main

import "fmt"

type Human interface {
	speak()
	walk()
}
type Dog struct {
	name string
}

func (d Dog) speak() string {
	return d.name + " is speaking"
}

func main() {
	d := Dog{
		name: "Border",
	}
	fmt.Println(d.speak())
}
