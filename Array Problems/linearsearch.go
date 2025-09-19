package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("enter the target u want to found: ")
	var x int
	fmt.Scan(&x)
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			fmt.Println("Target Found In the Array")
		}
	}

}
