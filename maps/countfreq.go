package main

import "fmt"

func main() {
	var str string
	fmt.Println("////////////////    Char Freq Counter////////////////////")
	fmt.Println("enter the string:")
	fmt.Scan(&str)
	runes := []rune(str)
	mp := make(map[rune]int)
	for i := 0; i < len(runes); i++ {
		mp[runes[i]]++
	}
	for k, v := range mp {
		fmt.Println(string(k), "its frequency is", v)
	}
}
