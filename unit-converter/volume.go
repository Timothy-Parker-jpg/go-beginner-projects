package main

var volumeAliases = map[string]string{
	"l":      "liter",
	"liter":  "liter",
	"liters": "liter",
	"litre":  "liter",
	"litres": "liter",

	"ml":          "milliliter",
	"milliliter":  "milliliter",
	"milliliters": "milliliter",
	"millilitre":  "milliliter",
	"millilitres": "milliliter",

	"gal":     "gallon",
	"gallon":  "gallon",
	"gallons": "gallon",

	"qt":     "quart",
	"quart":  "quart",
	"quarts": "quart",

	"pt":    "pint",
	"pint":  "pint",
	"pints": "pint",

	"cup":  "cup",
	"cups": "cup",

	"floz":        "floz",
	"fl oz":       "floz",
	"fluidounce":  "floz",
	"fluidounces": "floz",
}

var volumeToLiters = map[string]float64{
	// Metric
	"liter":      1,
	"milliliter": 0.001,
	"m3":         1000,  // cubic meter
	"cm3":        0.001, // cubic centimeter

	// US customary
	"tsp":    0.00492892,
	"tbsp":   0.0147868,
	"floz":   0.0295735,
	"cup":    0.236588,
	"pint":   0.473176,
	"quart":  0.946353,
	"gallon": 3.78541,

	// "gal_imp":  4.54609,
	// "qt_imp":   1.13652,
	// "pt_imp":   0.568261,
	// "floz_imp": 0.0284131,
}
