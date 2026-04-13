package main

var volumeAliases = map[string]string{
	"l":      "l",
	"liter":  "l",
	"liters": "l",
	"litre":  "l",
	"litres": "l",

	"ml":          "ml",
	"milliliter":  "ml",
	"milliliters": "ml",
	"millilitre":  "ml",
	"millilitres": "ml",

	"gal":     "gal",
	"gallon":  "gal",
	"gallons": "gal",

	"qt":     "qt",
	"quart":  "qt",
	"quarts": "qt",

	"pt":    "pt",
	"pint":  "pt",
	"pints": "pt",

	"cup":  "cup",
	"cups": "cup",

	"floz":        "floz",
	"fl oz":       "floz",
	"fluidounce":  "floz",
	"fluidounces": "floz",
}

var volumeToLiters = map[string]float64{
	// Metric
	"l":   1,
	"ml":  0.001,
	"m3":  1000,  // cubic meter
	"cm3": 0.001, // cubic centimeter

	// US customary
	"tsp":  0.00492892,
	"tbsp": 0.0147868,
	"floz": 0.0295735,
	"cup":  0.236588,
	"pt":   0.473176,
	"qt":   0.946353,
	"gal":  3.78541,

	// "gal_imp":  4.54609,
	// "qt_imp":   1.13652,
	// "pt_imp":   0.568261,
	// "floz_imp": 0.0284131,
}
