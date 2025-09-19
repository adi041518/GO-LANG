package main

import (
	"fmt"
)

func subtract(a int, b int) (int, error) {

	if a < b {
		err := fmt.Errorf("This cant be divisible %d is greater than %d", a, b)
		return 0, err
	}
	return a - b, nil
}

func main() {
	res, err := subtract(4, 5)
	if err != nil {
		fmt.Println("error generated", err)
	}
	fmt.Println(res)
}
