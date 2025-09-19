package main

import "fmt"

type person struct {
	id     int
	name   string
	gender string
	age    int
}

func main() {
	var p1 person
	fmt.Println("Enter the person id")
	fmt.Scan(&p1.id)
	fmt.Println("Enter the person name")
	fmt.Scan(&p1.name)
	fmt.Println("Enter the person gender")
	fmt.Scan(&p1.gender)
	fmt.Println("Enter the person age")
	fmt.Scan(&p1.age)
	fmt.Println(p1)

}
