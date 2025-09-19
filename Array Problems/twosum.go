package main

import "fmt"

func abs(a int, b int) int {
	if a < 0 {
		a = -(a)
	}
	if b < 0 {
		b = -(b)
	}
	var res int
	res = (a - b)
	return res
}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	s := make([][]int, 5)
	mp := make(map[int]int)
	var target int
	fmt.Scan(&target)
	for i := 0; i < len(arr); i++ {
		var x int

		key, _ := mp[arr[i]]
		x = abs(target, key)
		if x {
			newslice := []int{key, arr[i]}
			s = append(s, newslice)
		} else {

			fmt.Print("not found")
		}
	}
	fmt.Println(s)

}
