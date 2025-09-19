package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := os.Open("adi.txt")
	if err != nil {
		fmt.Println("error in file opening")
	}
	defer data.Close()

}
