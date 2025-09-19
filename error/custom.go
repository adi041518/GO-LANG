package main

import (
	"errors"
	"fmt"
)

func subtract(a int, b int) (int, error) {

	if a < b {
		err := errors.New("This cant be divisible")
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
