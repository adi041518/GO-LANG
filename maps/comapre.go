package main

import (
	"fmt"
	"reflect"
)

func main() {
	mp1 := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	mp2 := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	result := reflect.DeepEqual(mp1, mp2)
	if result {
		fmt.Println("both are equal")
	}

}
