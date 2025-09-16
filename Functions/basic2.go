package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	var x int
	var y int
	fmt.Println("Enter th First Value: ")
	fmt.Scan(&x)
	fmt.Println("Enter the Second Value: ")
	fmt.Scan(&y)
	fmt.Println(add(x, y))
}
