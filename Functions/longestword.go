package main

import "fmt"

func longestword(words ...string) string {
	var max int
	var word string
	for _, j := range words {
		x := len(j)
		if max < x {
			max = x
			word = j
		}
	}
	return word
}
func main() {
	s := longestword("Adtya", "works", "At ", "Tectoro")
	fmt.Println(s)
}
