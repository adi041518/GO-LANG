package main

import "fmt"

func main() {
	var a int
	fmt.Println("Divisiblity Checker of 3 and 5")
	fmt.Println("Enter the Number ")
	fmt.Scan(&a)
	switch {
	case a%3 == 0 && a%5 == 0:
		fmt.Println("Divisible by both 3 and 5")
		break

	case a%3 == 0:
		fmt.Println("Divisible by 3 Only")

	case a%5 == 0:
		fmt.Println("Divisible by 5 oNly")
		break
	default:
		fmt.Println("Not Divisible by Both 3 and 5")
		break

	}
}
