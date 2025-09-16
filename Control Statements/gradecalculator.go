package main

import "fmt"

func main() {
	fmt.Println("Enter the Marks You Received:")
	var a int
	fmt.Scan(&a)
	if (90 <= a) && (a <= 100) {
		fmt.Println("You Received A grade")
	} else if a >= 80 {
		fmt.Println("You Received B grade")
	} else if a >= 70 {
		fmt.Println("You Received C grade")
	} else {
		fmt.Println("You are Not Passed Sorry ( ;) )")
	}
}
