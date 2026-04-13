package main

func normalizeUnit(unit string, aliasMap map[string]string) (string, bool) {
	normalizedUnit, ok := aliasMap[unit]
	return normalizedUnit, ok
}

func convert(from, to *string, value *float64) (float64, error) {

	// var convertedValue float64
	// switch {
	// case IsLength(*from) && IsLength(*to):
	// 	convertedValue = ConvertLength(*value, *from, *to)
	// case IsWeight(*from) && IsWeight(*to):
	// 	convertedValue = ConvertWeight(*value, *from, *to)
	// default:
	// 	var zero float64
	// 	err := errors.New("Error: Both -from and -to flags must be a supported metrics.\nBoth metrics must be of the same type.\nUse -list flage to determine useable metrics.")
	// 	return zero, err
	// }
	// return convertedValue, nil
}
