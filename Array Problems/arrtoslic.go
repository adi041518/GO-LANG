package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]
	fmt.Println(arr)
	fmt.Println(slice)
}
