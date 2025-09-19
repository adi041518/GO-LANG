package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("das.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("there is no file present", err)
	}
}
