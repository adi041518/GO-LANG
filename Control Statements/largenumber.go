package main

import "fmt"

func main() {
	var x int
	var y int
	var z int
	fmt.Println("Enter the first Number: ")
	fmt.Scan(&x)
	fmt.Println("Enter the Second Number: ")
	fmt.Scan(&y)
	fmt.Println("Enter the Third Number: ")
	fmt.Scan(&z)
	if x > y {
		if x > z {
			fmt.Println(x, "is Greater Number")
		} else {
			fmt.Println(z, "is Greater Number")
		}
	} else if y > z {
		fmt.Println(y, "is Greater Number")
	} else {
		fmt.Println(x, "is Greater Number")
	}
}
