package string_calculator_go

import "testing"

func Test(t *testing.T) {
	sum := Calculator("")
	expected := 0
	if sum != expected {
		t.Errorf("Expected %q but got %q", expected, sum)
	}
}
