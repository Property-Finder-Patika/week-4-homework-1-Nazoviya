package main

import "fmt"

func main() {
	welcomeMessage :=
		`
-- Welcome to the temperature conversion program --
- type 'C' to convert for Celsius
- type 'F' to convert for Fahrenheit
- type 'K' to convert for Kelvin
`
	fmt.Println(welcomeMessage)
	unit, value := getTempType()
	tempConversion(unit, value)
}

// get user input and pass values.
func getTempType() (string, float64) {
	// get input type from user
	fmt.Print("Provide a unit to calculate: ")
	var unit string
	fmt.Scanln(&unit)
	// get input value from user
	fmt.Print("Provide a degree to convert: ")
	var value float64
	fmt.Scanln(&value)

	return unit, value
}

// comparision and call functions for temperature conversion.
func tempConversion(unit string, value float64) {
	switch unit {
	case "C":
		fmt.Printf("%v°C = %v°F and %v°K", value, CtoF(value), CtoK(value))
	case "F":
		fmt.Printf("%v°F = %v°C and %v°K", value, FtoC(value), FtoK(value))
	case "K":
		fmt.Printf("%v°K = %v°C and %v°F", value, KtoC(value), KtoF(value))
	default:
		fmt.Printf("%v is not available. Please use C, F, or K.", unit)
	}
}

// functions to calculate temperature conversions.
func CtoF(c float64) float64 {
	return c*1.8 + 32
}

func CtoK(c float64) float64 {
	return c + 273.15
}

func FtoC(f float64) float64 {
	return (f - 32) * .5556
}

func FtoK(f float64) float64 {
	return (f + 459.67) * 0.5556
}

func KtoC(k float64) float64 {
	return k - 273.15
}

func KtoF(k float64) float64 {
	return (k-273.15)*1.8 + 32
}
