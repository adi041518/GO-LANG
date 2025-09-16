package main

import (
	"fmt"

	"math"
)

func main() {
	fmt.Println("Enter the Armstrong Number: ")
	var x int
	fmt.Scan(&x)
	var y = x
	var sum int
	for x > 0 {
		n := x % 10
		sum += int(math.Pow(float64(n), 3))
		x /= 10
	}
	if y == sum {
		fmt.Println(y, "is an Armstrong Number")
	}
}
