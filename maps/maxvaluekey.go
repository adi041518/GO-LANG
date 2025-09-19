package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("////////////////    Char Freq Counter////////////////////")
	fmt.Println("enter the string:")
	var str string
	fmt.Scan(&str)
	str = strings.ToLower(str)
	runes := []rune(str)
	mp := make(map[rune]int)
	for i := 0; i < len(runes); i++ {
		mp[runes[i]]++
	}
	var maxv int

	for _, v := range mp {
		if maxv < v {
			maxv = v
		}
	}
	fmt.Println("Maximum Value is: ", maxv)
}
