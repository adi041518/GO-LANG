package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("Adi.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("error in opening/creating a file")
		return
	}
	defer file.Close()
	lines := 0
	words := 0
	chars := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		lines++
		chars += len(line)
		words += len((strings.Fields(line)))
	}
	fmt.Println(words)
	fmt.Println(chars)
}
