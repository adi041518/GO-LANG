package main

import "fmt"

func main() {
	a, b := 4, 5
	fmt.Println("Before Swapping numbers are", a, b)

	b, a = a, b
	fmt.Println("After Swapping numbers are", a, b)

}
