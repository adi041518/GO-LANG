package main

import "fmt"

func main() {
	c := complex(3, 6)
	c1 := complex(4, 2)
	fmt.Println("Addition of two complex number is:")
	fmt.Println(c * c1)
	fmt.Println("Addition of Real numbers is:")
	fmt.Println(real(c) + real(c1))
	fmt.Println("Addition of imaginary numbers are: ")
	fmt.Println(imag(c) + imag(c1))
}
