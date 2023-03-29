package conv

import "math"

/*
	FahrenheitToCelsius
	CelsiusToFahrenheit
	KelvinToFahrenheit
	CelsiusToKelvin
	KelvinToCelsius
	FahrenheitToKelvin
*/

// La til math.Round slik at output bare har 2 tall bak komma.
func FahrenheitToCelsius(Fahrenheit float64) float64 {
	Celsius := (Fahrenheit - 32) * 5 / 9
	return math.Round(Celsius*100) / 100
}

func CelsiusToFahrenheit(Celsius float64) float64 {
	Fahrenheit := (Celsius * 1.8) + 32
	return math.Round(Fahrenheit*100) / 100
}

func KelvinToFahrenheit(Kelvin float64) float64 {
	Fahrenheit := Kelvin*9/5 - 459.67
	return math.Round(Fahrenheit*100) / 100
}

func CelsiusToKelvin(Celsius float64) float64 {
	Kelvin := Celsius + 273.15
	return math.Round(Kelvin*100) / 100
}

func KelvinToCelsius(Kelvin float64) float64 {
	Celsius := Kelvin - 273.15
	return math.Round(Celsius*100) / 100

}
func FahrenheitToKelvin(Fahrenheit float64) float64 {
	Kelvin := (Fahrenheit + 459.67) * 5 / 9
	return math.Round(Kelvin*100) / 100
}
