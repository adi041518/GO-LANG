package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	x := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = x
	fmt.Println(arr)
}
