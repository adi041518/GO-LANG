package main

import (
	"fmt"
)

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := len(arr)
	fmt.Println("Enter the Target u want:")
	var x int
	fmt.Scan(&x)
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == x {
			fmt.Println("target Found")
			break
		}
		if x < arr[mid] {
			high = mid - 1
		}
		if x > arr[mid] {
			low = mid + 1
		}
	}
}
