package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Person struct {
	firstname string
	lastname  string
	A         Address
}

func main() {

	p := Person{
		firstname: "Aditya",
		lastname:  "Kadambala",
		A: Address{
			City:  "Srikakulam",
			State: "AndhraPradesh",
		},
	}
	fmt.Println(p.A.City)

}
