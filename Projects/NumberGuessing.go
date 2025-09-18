package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(100)
	fmt.Println("Enter the number 1 to 100 randomly")
	var b int
	fmt.Scan(&b)
	if b == x {
		fmt.Println("you are right")
	} else {
		fmt.Println("OOPS sorry next time")
	}

}
