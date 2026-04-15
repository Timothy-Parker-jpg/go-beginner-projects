package main

import (
	"fmt"
)

type UnitType string

const (
	Length UnitType = "length"
	Weight UnitType = "weight"
	Temp   UnitType = "temp"
	Volume UnitType = "volumne"
)

type UnitDef struct {
	Unit string
	Type UnitType
	// Factor float64
}

func initMasterAliasMap() map[string]UnitDef {
	MasterAliasMap := make(map[string]UnitDef)
	for k, v := range lengthAliases {
		MasterAliasMap[k] = UnitDef{Unit: v, Type: Length}
	}
	for k, v := range weightAliases {
		MasterAliasMap[k] = UnitDef{Unit: v, Type: Weight}
	}
	for k, v := range tempAliases {
		MasterAliasMap[k] = UnitDef{Unit: v, Type: Temp}
	}
	for k, v := range volumeAliases {
		MasterAliasMap[k] = UnitDef{Unit: v, Type: Volume}
	}

	return MasterAliasMap
}

func listUnits() {
	fmt.Println("\t\tSupported Units")
	fmt.Println("Length:")
	for k, _ := range lengthToMeter {
		fmt.Println("\t" + k)
	}
	fmt.Println("Weight:")
	for k, _ := range weightToKilogram {
		fmt.Println("\t" + k)
	}
	fmt.Println("Temp:")
	for k, _ := range tempAliases {
		fmt.Println("\t" + k)
	}
	fmt.Println("Volume:")
	for k, _ := range volumeToLiters {
		fmt.Println("\t" + k)
	}

}
