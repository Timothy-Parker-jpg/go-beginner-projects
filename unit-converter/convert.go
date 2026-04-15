package main

import (
	"errors"
	"fmt"
)

func normalizeUnit(unit string, aliasMap map[string]UnitDef) (UnitDef, bool) {
	normalizedUnit, ok := aliasMap[unit]
	return normalizedUnit, ok
}

func convert(from, to *string, value *float64) error {
	masterAlias := initMasterAliasMap()
	normFrom, ok := normalizeUnit(*from, masterAlias)
	if !ok {
		// var zero float64
		return errors.New("Error: failed to normalize -from.")
	}
	normTo, ok := normalizeUnit(*to, masterAlias)

	if !ok {
		// var zero float64
		return errors.New("Error: failed to normalize -to.")
	}

	if normTo.Type == normFrom.Type {
		switch normTo.Type {
		case Length:
			format := "%.2f %s = %.2f %s\n"
			convertedValue := *value * lengthToMeter[normFrom.Unit] / lengthToMeter[normTo.Unit]
			fmt.Printf(format, *value, normFrom, convertedValue, normTo)
			return nil
		case Weight:
			format := "%.2f %s = %.2f %s\n"
			convertedValue := *value * weightToKilogram[normFrom.Unit] / weightToKilogram[normTo.Unit]
			fmt.Printf(format, *value, normFrom, convertedValue, normTo)
			return nil
		case Temp:
			var c float64
			var f float64
			var k float64
			switch normFrom.Unit {
			case "celsius":
				c = *value
				f = CelsiusToFahrenheit(*value)
				k = CelsiusToKelvin(*value)
			case "fahrenheit":
				c = FahrenheitToCelsius(*value)
				f = *value
				k = FahrenheitToKelvin(*value)
			case "kelvin":
				c = KelvinToCelsius(*value)
				f = KelvinToFahrenheit(*value)
				k = *value
			}
			format := "%.2f°C = %.2f°F = %.2fK\n"
			fmt.Printf(format, c, f, k)
			return nil
		case Volume:
			format := "%.2f %s = %.2f %s"
			convertedValue := *value * volumeToLiters[normFrom.Unit] / volumeToLiters[normTo.Unit]
			fmt.Printf(format, *value, normFrom, convertedValue, normTo)
			return nil

		}
	}
	return errors.New("Error: -from & -to Types do not match.")
}
