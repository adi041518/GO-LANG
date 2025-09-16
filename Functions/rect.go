package main

import "fmt"

func rect(a int, b int) (area int, perimeter int) {
	area = a * b
	perimeter = 2 * (a + b)
	return
}

func main() {
	a, b := rect(5, 4)
	fmt.Println(a, b)
}
