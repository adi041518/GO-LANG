package main

import "fmt"

func main() {
	mp := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	for k, v := range mp {
		fmt.Println(k, " ", v)
	}
}
