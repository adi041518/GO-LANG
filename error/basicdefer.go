package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("das.txt")
	if err != nil {
		fmt.Println("error is existed")
	}
	defer file.Close()
}
