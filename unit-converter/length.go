package main

import (
	"sort"
)

var lengthAliases = map[string]string{
	"m":      "meter",
	"meter":  "meter",
	"meters": "meter",

	"km":         "kilometer",
	"kilometer":  "kilometer",
	"kilometers": "kilometer",

	"cm":          "centimeter",
	"centimeter":  "centimeter",
	"centimeters": "centimeter",

	"mm":          "millimeter",
	"millimeter":  "millimeter",
	"millimeters": "millimeter",

	"in":     "inch",
	"inch":   "inch",
	"inches": "inch",

	"ft":   "foot",
	"foot": "foot",
	"feet": "foot",

	"yd":    "yard",
	"yard":  "yard",
	"yards": "yard",

	"mi":    "mile",
	"mile":  "mile",
	"miles": "mile",
}

var lengthMap = map[string]float64{
	"millimeter": 0.001,
	"centimeter": 0.01,
	"meter":      1,
	"kilometer":  1000,
	"inch":       0.0254,
	"foot":       0.3048,
	"mile":       1609.34,
}

func listLengths() []string {
	keys := make([]string, 0, len(lengthMap))

	for k := range lengthMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys

}

func IsLength(unit string) bool {
	if _, ok := lengthAliases[unit]; ok {
		return true
	}
	return false
}

func ConvertLength(value float64, from, to string) float64 {
	return value * lengthMap[from] / lengthMap[to]
}
