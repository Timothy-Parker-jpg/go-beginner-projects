package main

import (
	"sort"
)

var weightAliases = map[string]string{
	"g":     "g",
	"gram":  "g",
	"grams": "g",

	"kg":        "kg",
	"kilogram":  "kg",
	"kilograms": "kg",

	"mg":         "mg",
	"milligram":  "mg",
	"milligrams": "mg",

	"lb":     "lb",
	"pound":  "lb",
	"pounds": "lb",

	"oz":     "oz",
	"ounce":  "oz",
	"ounces": "oz",
}

var weightMap = map[string]float64{
	"mg":  0.000001,  // milligram
	"g":   0.001,     // gram
	"kg":  1,         // base unit
	"lb":  0.453592,  // pound
	"oz":  0.0283495, // ounce
	"ton": 1000,      // metric ton
}

func listWeights() []string {
	keys := make([]string, 0, len(weightMap))

	for k := range weightAliases {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func IsWeight(unit string) bool {
	if _, ok := weightMap[unit]; ok {
		return true
	}
	return false
}
func ConvertWeight(value float64, from, to string) float64 {
	return value * weightMap[from] / weightMap[to]
}
