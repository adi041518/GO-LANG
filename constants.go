package main

import (
	"fmt"
)

func main() {
	const pi = 3.14
	var area float64
	fmt.Println("enter the Radius:")
	fmt.Scan(&area)
	x := pi * area
	fmt.Println("area of circle is", x)
}
