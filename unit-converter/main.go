package main

import (
	"flag"
	"fmt"
	"os"
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

	// MasterAliasMap := initMasterAliasMap()

	if *list == true {
		listUnits()
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
	err := convert(from, to, value)
	fmt.Println("Conversion Complete")
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR: Convert() : %w", err))
	}
	os.Exit(1)
}
