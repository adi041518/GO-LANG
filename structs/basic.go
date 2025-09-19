package main

import "fmt"

type Book struct {
	title  string
	author string
	page   int
}

func main() {
	var b1 Book
	var b2 = Book{
		title:  "bannana",
		author: "Aditya",
		page:   123,
	}
	fmt.Println(b1, b2)
}
