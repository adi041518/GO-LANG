package main

import (
	"fmt"
	"math"
)

func main() {
	c := func(x int) int {
		return int(math.Pow(float64(x), 3))
	}
	fmt.Println(c(3))

}
