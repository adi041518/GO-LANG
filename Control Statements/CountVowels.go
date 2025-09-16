package main

import "fmt"

func main() {
	str := "geeks for geeks"
	var count int
	for _, j := range str {
		if j == 'a' || j == 'e' || j == 'i' || j == 'e' || j == 'o' || j == 'u' {
			count++
		}
	}
	fmt.Println("the number of vowels in the string are: ", count)
}
