package string_calculator_go

import "strconv"

func Calculator(s string) int {
	if s == "" {
		return 0
	}
	i, _ := strconv.Atoi(s)
	return i
}
