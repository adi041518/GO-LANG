package main

import "fmt"

func add(a *int, b *int) int {
	return *a + *b
}
func main() {
	var x int
	var y int
	fmt.Println("enter the first value: ")
	fmt.Scan(&x)
	fmt.Println("enter the second value: ")
	fmt.Scan(&y)
	c := add(&x, &y)
	fmt.Println(c)
}
