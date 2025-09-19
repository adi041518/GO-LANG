package main

import "fmt"

type Complex struct {
	real      int
	imaginary int
}

func main() {
	C := Complex{3, 10}
	p := &C
	p.real = 15
	fmt.Println(C)
}
