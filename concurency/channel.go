package main

import "fmt"

func say(ch chan int, b int) {

	ch <- b
}
func main() {
	ch := make(chan int)
	go say(ch, 101)
	var x int
	x = <-ch
	fmt.Println(x)
}
