package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	for _, v := range s {
		fmt.Println(v)
	}
	//slice from array
	arr := [5]int{10, 20, 30, 40, 50}
	s1 := arr[0:2]
	fmt.Println(s1)
	//appending

	var s2 []int
	s2 = append(s2, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(s2)

	//copy function

	s3 := []int{1, 2, 3, 4, 5}
	s4 := slices.Clone(s3)
	fmt.Println(s4)
	res := reflect.DeepEqual(s3, s4)
	if res {
		fmt.Println("both are equal")
	}

	//to say slice is referece type

	arr1 := [5]int{1, 2, 3, 4, 5}
	s5 := arr1[0:]
	s5[0] = 22
	fmt.Println(s5, arr1)

}
