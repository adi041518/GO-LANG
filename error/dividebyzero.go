package main

import "fmt"

func main() {
	a := 4
	b := 5
	defer func() {
		if res := recover(); res != nil {
			fmt.Println("recovered from panic")
		}
	}()
	if a < b {
		panic("Error Occured")
	}

}
