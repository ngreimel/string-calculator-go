package string_calculator_go

import (
	"strconv"
	"strings"
)

func Calculator(valuesToAdd string) int {
	if valuesToAdd == "" {
		return 0
	}
	f := func(c rune) bool {
		return c == ',' || c == '\n'
	}
	stringsSplitByDelimiter := strings.FieldsFunc(valuesToAdd, f)
	return AddValues(stringsSplitByDelimiter)
}

func AddValues(stringsSplitByDelimiter []string) int {
	sum := 0
	for i := 0; i < len(stringsSplitByDelimiter); i++ {
		intValue, _ := strconv.Atoi(stringsSplitByDelimiter[i])
		sum += intValue
	}
	return sum
}
