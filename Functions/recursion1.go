package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	sum := fact(5)
	fmt.Println("The factorial of the numebr is ", sum)
}
