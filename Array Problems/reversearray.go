package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	n := len(arr)
	for i := 0; i < len(arr); i++ {
		if i <= n-i-1 {
			arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
		}
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
