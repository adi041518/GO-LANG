package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Println("Enter the String: ")
	fmt.Scan(&x)
	fmt.Println("Lenght of the string is: ", len(x))
	var vcount int
	var ccount int
	x = strings.ToLower(x)
	for _, v := range x {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			vcount++
		} else {
			ccount++
		}
	}
	fmt.Println("Number of Vowels: ", vcount)
	fmt.Println("Number of Consonants: ", ccount)
	i := 0
	j := len(x) - 1
	var boo bool = true
	for i <= j {
		if x[i] != x[j] {
			boo = false
		}
		i++
		j--
	}

	if boo {
		fmt.Println("YAy this is a palindrome")
	}

}
