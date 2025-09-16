package main

import "fmt"

func mul(nums ...int) int {
	var mulsum = 1
	for _, j := range nums {
		mulsum *= j
	}
	return mulsum
}

func main() {
	fmt.Println(mul(1, 2, 3))
	fmt.Println(mul(1, 2))
}
