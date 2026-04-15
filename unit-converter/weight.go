package main

import (
	"sort"
)

var weightAliases = map[string]string{
	"g":     "gram",
	"gram":  "gram",
	"grams": "gram",

	"kg":        "kilogram",
	"kilogram":  "kilogram",
	"kilograms": "kilogram",

	"mg":         "milligram",
	"milligram":  "milligram",
	"milligrams": "milligram",

	"lb":     "pound",
	"pound":  "pound",
	"pounds": "pound",

	"oz":     "ounce",
	"ounce":  "ounce",
	"ounces": "ounce",
}

var weightToKilogram = map[string]float64{
	"milligram": 0.000001,  // milligram
	"gram":      0.001,     // gram
	"kilogram":  1,         // base unit
	"pound":     0.453592,  // pound
	"ounce":     0.0283495, // ounce
	"ton":       1000,      // metric ton
}

func listWeights() []string {
	keys := make([]string, 0, len(weightToKilogram))

	for k := range weightAliases {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func IsWeight(unit string) bool {
	if _, ok := weightToKilogram[unit]; ok {
		return true
	}
	return false
}
func ConvertWeight(value float64, from, to string) float64 {
	return value * weightToKilogram[from] / weightToKilogram[to]
}
