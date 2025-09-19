package main

import "fmt"

func main() {
	defer fmt.Println("tectoroo")
	defer fmt.Println("from")
	defer fmt.Println("Aditya")
	panic("this is Aditya")
	//defer is LIFO
}
