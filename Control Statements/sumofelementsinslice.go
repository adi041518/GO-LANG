package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	var sum int
	for _, j := range s {
		sum += j
	}
	fmt.Println("Sum of Elemenst in slice are: ", sum)
}
