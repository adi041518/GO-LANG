package main

import "fmt"

func cal(a int, b int, c int) (int, error) {
	switch c {
	case 1:
		return a + b, nil
	case 2:
		if a < b {
			err := fmt.Errorf("This cant be divisible")
			return 0, err
		}
		return a - b, nil
	case 3:
		return a * b, nil
	case 4:
		if a < b {
			fmt.Errorf("this cant be divisible because %d is less than %d", a, b)
		}
		return a / b, nil
	}
	return 0, nil
}

func main() {
	fmt.Println("Calculator of two numbers: ")
	var x int
	var y int
	var c int
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&c)
	res, err := cal(x, y, c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
