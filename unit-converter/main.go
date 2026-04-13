package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	// "errors"
)

func main() {
	fmt.Println("************** Unit Converter **************")
	// fmt.Print(">> ")
	value := flag.Float64("value", 0.0, "Input value")
	from := flag.String("from", "", "Source unit")
	to := flag.String("to", "", "To unit")
	list := flag.Bool("list", false, "list of available units")

	flag.Parse()

	initMasterAliasMap()

	if *list == true {
		fmt.Printf("Supported length units: %v\n", listLengths())
		fmt.Printf("Supported weight units: %v\n", listWeights())
		os.Exit(0)
	}

	if *value <= 0 {
		fmt.Println("Error: -value must use a value greater than 1.")
		os.Exit(1)
	}
	if *from == "" || *to == "" {
		fmt.Println("Error: -from and -to must contain a unit of measure.")
		os.Exit(1)
	}

	*from = strings.ToLower(*from)
	*to = strings.ToLower(*to)

	fmt.Printf("Processing %v from %v to %v ...\n", *value, *from, *to)

	var convertedValue float64

	switch {
	case IsLength(*from) && IsLength(*to):
		convertedValue = ConvertLength(*value, *from, *to)
	case IsWeight(*from) && IsWeight(*to):
		convertedValue = ConvertWeight(*value, *from, *to)
	default:
		fmt.Println("Error: Both -from and -to flags must be a supported length or weight metric.\nWeight and Length metics cannot be mixed.")
		os.Exit(1)
	}

	format := "%.2f %s = %.2f %s"

	fmt.Printf(format, *value, *from, convertedValue, *to)
}
