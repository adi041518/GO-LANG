package main

import (
	"fmt"
)

const (
	cToFMultiplier = 9.0 / 5.0
	cToFOffset     = 32.0
)

func celsiusToFahrenheit(c float64) float64 {
	return c*cToFMultiplier + cToFOffset
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - cToFOffset) / cToFMultiplier
}

func main() {
	var choice int
	var temp float64

	fmt.Println("=== Temperature Converter ===")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Print("Enter your choice (1/2): ")
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Print("Enter temperature in Celsius: ")
		fmt.Scanln(&temp)
		fmt.Printf("%.2f °C = %.2f °F\n", temp, celsiusToFahrenheit(temp))
	} else if choice == 2 {
		fmt.Print("Enter temperature in Fahrenheit: ")
		fmt.Scanln(&temp)
		fmt.Printf("%.2f °F = %.2f °C\n", temp, fahrenheitToCelsius(temp))
	} else {
		fmt.Println("Invalid choice!")
	}
}
