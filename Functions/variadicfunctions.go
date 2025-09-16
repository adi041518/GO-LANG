package main

import "fmt"

func add(nums ...int) int {
	var sum int
	for _, j := range nums {
		sum += j
	}
	return sum
}

func main() {

	fmt.Println(add(1, 2, 3))
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3, 4))
}
