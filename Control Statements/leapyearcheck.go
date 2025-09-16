package main

import "fmt"

func main() {
	var x int
	fmt.Println("Enter the year ;")
	fmt.Scan(&x)
	if (x%100 != 0 && x%4 == 0) || x%400 == 0 {
		fmt.Println(x, "is a leap year")
	} else {
		fmt.Println(x, "is not a leap year")
	}
}
