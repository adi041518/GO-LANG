package main

import "fmt"

func Square(a int) int {
	return a * a
}

func main() {
	fmt.Println("enter the number u want: ")
	var a int
	fmt.Scan(&a)
	fmt.Println("The Square of the Number is : ", Square(a))
}
