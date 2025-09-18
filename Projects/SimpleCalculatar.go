package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Println("enter the two numbers: ")
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Println("Addition of two numbers are: ", x+y)
	fmt.Println("Subtraction of two numbers are: ", x-y)
	fmt.Println("Multiplication of two numbers are: ", x*y)
	fmt.Println("Division of two numbers are: ", x/y)
}
