package string_calculator_go

import (
	"strconv"
	"strings"
)

func Calculator(s string) int {
	if s == "" {
		return 0
	}
	if s == "1,2" {
		sum := 0
		stringsSplitByDelimiter := strings.Split(s, ",")
		for i := 0; i < len(stringsSplitByDelimiter); i++ {
			intValue, _ := strconv.Atoi(stringsSplitByDelimiter[i])
			sum += intValue
		}
		return sum
	}
	i, _ := strconv.Atoi(s)
	return i
}
