package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{4, 5, 6, 1, 2, 3}
	sort.Ints(s)
	fmt.Println(s)
}
