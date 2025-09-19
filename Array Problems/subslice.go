package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[1:5]
	s2 := s[5:8]
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
}
