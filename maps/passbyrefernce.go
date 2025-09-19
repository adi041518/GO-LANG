package main

import "fmt"

func modi(m map[string]int) {
	m["E"] = 5
}

func main() {
	mp := make(map[string]int)
	mp["A"] = 1
	mp["B"] = 2
	mp["C"] = 5
	mp["D"] = 4
	mp["E"] = 3
	modi(mp)
	fmt.Println(mp)
}
