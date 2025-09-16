package main

import "fmt"

func count() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

func main() {
	c := count()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

}
