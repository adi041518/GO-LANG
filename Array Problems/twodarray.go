package main

import "fmt"

func main() {
	arr := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
