package main

import "fmt"

func main() {
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var dsum int
	var esum int
	var n = len(matrix)
	for i := 0; i < len(matrix); i++ {
		dsum += matrix[i][n-i-1]
		esum += matrix[i][i]
	}
	fmt.Println(dsum + esum)
}
