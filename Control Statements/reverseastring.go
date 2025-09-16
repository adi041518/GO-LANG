package main

import "fmt"

func main() {
	var s string
	fmt.Println("enter the string u want to reverse it: ")
	fmt.Scan(&s)
	runes := []rune(s)
	i := 0
	j := len(runes) - 1
	for i <= j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	fmt.Println("reversed string is: ")
	for _, j := range runes {
		fmt.Print(string(j))
	}
	fmt.Println()
}
