package string_calculator_go

import "testing"

func Test(t *testing.T) {
	sum := Calculator("")
	expected := 0
	if sum != expected {
		t.Errorf("Expected %q but got %q", expected, sum)
	}
}

func Test2(t *testing.T) {
	sum := Calculator("1")
	expected := 1
	if sum != expected {
		t.Errorf("Expected %q but got %q", expected, sum)
	}
}

func Calculator(s string) interface{} {
	return 0
}
