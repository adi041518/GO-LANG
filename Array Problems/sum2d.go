package main

import "fmt"

func main() {
	var sum int
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			sum += matrix[i][j]
		}
	}
	fmt.Println(sum)
}
