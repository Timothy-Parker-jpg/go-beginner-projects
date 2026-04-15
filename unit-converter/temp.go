package main

import (
	"errors"
)

var tempAliases = map[string]string{
	"c":       "celsius",
	"celsius": "celsius",
	"degc":    "celsius",

	"f":          "fahrenheit",
	"fahrenheit": "fahrenheit",
	"degf":       "fahrenheit",

	"k":      "kelvin",
	"kelvin": "kelvin",
}

func CelsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}
func CelsiusToKelvin(c float64) float64 {
	return c + 273.15
}
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}
func FahrenheitToKelvin(f float64) float64 {
	return (f-32)*5/9 + 273.15
}
func KelvinToFahrenheit(k float64) float64 {
	return (k-273.15)*9/5 + 32
}
func KelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func TempConvert(value float64, unit string) (c float64, f float64, k float64, err error) {

	switch unit {
	case "celsius":
		c = value
		f = CelsiusToFahrenheit(value)
		k = CelsiusToKelvin(value)
	case "fahrenheit":
		c = FahrenheitToCelsius(value)
		f = value
		k = FahrenheitToKelvin(value)
	case "kelvin":
		c = KelvinToCelsius(value)
		f = KelvinToFahrenheit(value)
		k = value
	default:
		err = errors.New("Unsupported unit. Please use C, F, K (Celsius, Fahrenheit, Kelvin)")
	}
	return c, f, k, err
}
