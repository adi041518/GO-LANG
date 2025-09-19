package main

import (
	"fmt"
	"time"
)

func firstgoroutine() {
	for i := 0; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("hii")
	}
}
func secondgoroutine() {
	for i := 0; i < 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Println("hiiii")
	}
}
func main() {
	go firstgoroutine()
	go secondgoroutine()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("hello")
}
