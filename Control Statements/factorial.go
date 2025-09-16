package main

import "fmt"

func main() {
	fmt.Println("enter the Number:")
	var a int
	fmt.Scan(&a)
	sum := 1
	for i := 1; i <= a; i++ {
		sum *= i
	}
	fmt.Printf("factorial of %d: %d", a, sum)
	fmt.Println()
}
