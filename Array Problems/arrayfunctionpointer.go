package main

import "fmt"

func add(ar *[5]int) int {
	var sum int
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}
	ar[4] = 21
	return sum
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	sum := add(&arr)
	fmt.Println(sum)
	fmt.Println(arr)
}
