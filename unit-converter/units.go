package main

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

var MasterAliasMap map[string]UnitDef

func initMasterAliasMap() {
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
}

func listUnits() {
	values := make([]UnitDef, 0, len(MasterAliasMap))
	sort.
}
