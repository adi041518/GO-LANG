package main

import "fmt"

func main() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} //2d slice
	for i, v := range matrix {
		for j, k := range v {
			fmt.Println(matrix[i][j], k)
		}
	}
}
