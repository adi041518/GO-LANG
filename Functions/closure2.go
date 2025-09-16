package main

import (
	"fmt"
	"math"
)

func outer(n int) func(float64) float64 {
	return func(x float64) float64 {
		return math.Pow(x, float64(n))
	}
}

func main() {
	square := outer(2)
	result := square(5)
	fmt.Println(result)

}
