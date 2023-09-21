package string_calculator_go

import (
	"strconv"
	"strings"
)

func Calculator(valuesToAdd string) int {
	if valuesToAdd == "" {
		return 0
	}
	return AddValues(valuesToAdd)
}

func AddValues(valuesToAdd string) int {
	sum := 0
	stringsSplitByDelimiter := strings.Split(valuesToAdd, ",")
	for i := 0; i < len(stringsSplitByDelimiter); i++ {
		intValue, _ := strconv.Atoi(stringsSplitByDelimiter[i])
		sum += intValue
	}
	return sum
}
