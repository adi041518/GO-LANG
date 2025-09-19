package main

import "fmt"

func main() {
	//from here we can notice that arrays are value type not reference type
	arr := [6]int{1, 2, 3, 4, 5}
	arr1 := arr
	arr[2] = 5
	fmt.Println(arr1)
	fmt.Println(arr)

}
