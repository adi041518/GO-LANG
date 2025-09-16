package main

import "fmt"

func main() {
	fmt.Println("Enter the Age to vote:")
	var x int
	fmt.Scan(&x)
	if x > 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("Sorry you are under age! please try again after", 18-x, "years")
	}
}
