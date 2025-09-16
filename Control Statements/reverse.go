package main

import "fmt"

func main() {
	var x int
	fmt.Println("enter the Number:")
	fmt.Scan(&x)
	var rev int
	for x > 0 {
		n := x % 10
		rev = rev*10 + n
		x /= 10

	}
	fmt.Println("reversed Number is: ", rev)

}
