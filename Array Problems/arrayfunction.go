package main

import "fmt"

func add(arr [5]int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	sum := add(arr)
	fmt.Println(sum)
}
