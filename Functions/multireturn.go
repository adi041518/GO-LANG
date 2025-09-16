package main

import "fmt"

func cal(a int, b int) (int, int, int) {
	return a + b, a - b, a * b
}

func main() {
	var x int
	var y int
	fmt.Println("Enter the First Value: ")
	fmt.Scan(&x)
	fmt.Println("Enter the Second Value: ")
	fmt.Scan(&y)
	var m, n, c = cal(x, y)
	fmt.Println(m, n, c)
}
