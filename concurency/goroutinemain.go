package main

import "fmt"

func say() {
	fmt.Println("this is not working")
}
func main() {
	go say()
	fmt.Println("this is working")
}

// this happens because when the goroutines call returns immediately this make it to not execute..
