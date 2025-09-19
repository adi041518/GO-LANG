package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	var max int
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Println(max)

}
