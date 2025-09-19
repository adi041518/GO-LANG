package main

import (
	"fmt"
	"time"
)

func routine() {
	time.Sleep(10000 * time.Millisecond)
	fmt.Println("hi this is Aditya")
}

func main() {
	go routine()
	routine()

}
