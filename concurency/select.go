package main

import "fmt"

func goOne(ch chan int) {
	ch <- 4
}

func main() {
	ch := make(chan int)
	dh := make(chan int)
	go goOne(ch)
	go goOne(dh)
	select {
	case x := <-ch:
		fmt.Println(x)
	case y := <-ch:
		fmt.Println(y)
	}

}
