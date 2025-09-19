package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}
type Dog struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}
func (d Dog) Speak() string {
	return "Hello, my name is " + d.Name
}

func main() {
	var s Speaker
	s = Person{Name: "Alice"}
	fmt.Println(s.Speak())
	value, ok := s.(Person)
	if ok {
		fmt.Println("true")
	}
	fmt.Println(value)
	var s1 Speaker
	s1 = Dog{
		Name: "Border",
	}
	fmt.Println(s1.Speak())
}
