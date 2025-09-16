package main

import "fmt"

func main() {
	c := func(a int) int {
		return a * a
	}
	var x int
	fmt.Println("Enter the Range of The Numbers U want : ")
	fmt.Scan(&x)
	fmt.Println("the Square are :")
	for i := 1; i <= x; i++ {
		fmt.Println(c(i))
	}
}
