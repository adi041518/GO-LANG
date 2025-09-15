package main

import "fmt"

func main() {
	var x int
	fmt.Println("enter the first number")
	fmt.Scan(&x)
	var y int
	fmt.Println("enter the second number")
	fmt.Scan(&y)

	sum := x + y
	fmt.Println("Sum of two numbers are ", sum)
}
