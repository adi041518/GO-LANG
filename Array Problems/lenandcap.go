package main

import "fmt"

func main() {
	//len and cap

	s := make([]int, 6, 8)
	s = append(s, 1)
	fmt.Println(cap(s))
	s = append(s, 2)
	fmt.Println(cap(s))
	s = append(s, 3)
	fmt.Println(cap(s))
	s = append(s, 4)
	fmt.Println(cap(s))
	s = append(s, 5)
	s = append(s, 6)
	fmt.Println(s)
}
