package main

import "fmt"

func main() {
	mp := map[string]int{
		"tv":             1,
		"washingmachine": 4,
		"fridge":         5,
	}
	fmt.Println("Enter the key u want to delete: ")
	var x string
	fmt.Scan(&x)
	delete(mp, x)
	fmt.Println(mp)
}
