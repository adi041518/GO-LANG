package main

import "fmt"

func is_Positive(a int) bool {
	if a < 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println("enter the number: ")
	var a int
	fmt.Scan(&a)
	fmt.Println(is_Positive(a))

}
