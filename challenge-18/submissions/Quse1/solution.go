package main

import (
	"fmt"
	"math"
)

func main() {
	// Example usage
	var celsius, fahrenheit float64

	_, err := fmt.Scan("%.f", celsius)
	if err != nil {
		fmt.Printf("Scan error standart input: %s\n", err)
		return
	}

	fahrenheit = CelsiusToFahrenheit(celsius)
	//fmt.Printf("%.2f°C is equal to %.2f°F\n", celsius, fahrenheit)
	fmt.Println(fahrenheit)

	_, err = fmt.Scan("%.f", fahrenheit)
	if err != nil {
		fmt.Printf("Scan error standart input: %s\n", err)
		return
	}

	celsius = FahrenheitToCelsius(fahrenheit)
	//fmt.Printf("%.2f°F is equal to %.2f°C\n", fahrenheit, celsius)
	fmt.Println(celsius)
}

// CelsiusToFahrenheit converts a temperature from Celsius to Fahrenheit
// Formula: F = C × 9/5 + 32
func CelsiusToFahrenheit(celsius float64) float64 {
	result := celsius*9/5 + 32
	// TODO: Implement this function
	// Remember to round to 2 decimal places
	return Round(result, 2)
}

// FahrenheitToCelsius converts a temperature from Fahrenheit to Celsius
// Formula: C = (F - 32) × 5/9
func FahrenheitToCelsius(fahrenheit float64) float64 {
	result := (fahrenheit - 32) * 5 / 9
	// TODO: Implement this function
	// Remember to round to 2 decimal places
	return Round(result, 2)
}

// Round rounds a float64 value to the specified number of decimal places
func Round(value float64, decimals int) float64 {
	precision := math.Pow10(decimals)
	return math.Round(value*precision) / precision
}
