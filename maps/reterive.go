package main

import "fmt"

func main() {
	mp := map[string]int{
		"tv":             1,
		"washingmachine": 4,
		"fridge":         5,
	}
	target := "fridge"
	valeu, exists := mp[target]
	if exists {
		fmt.Println("Tareger found", valeu)
	} else {
		fmt.Println("Target not found")
	}
}
