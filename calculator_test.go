package string_calculator_go

import "testing"

func Test(t *testing.T) {
	sum := Calculator("")
	expected := 0
	assertEqual(t, sum, expected)
}

func Test2(t *testing.T) {
	sum := Calculator("1")
	expected := 1
	assertEqual(t, sum, expected)
}

func Test3(t *testing.T) {
	sum := Calculator("2")
	expected := 2
	assertEqual(t, sum, expected)
}

func Test4(t *testing.T) {
	sum := Calculator("1,2")
	expected := 3
	assertEqual(t, sum, expected)
}

func assertEqual(t *testing.T, sum int, expected int) {
	if sum != expected {
		t.Errorf("Expected %q but got %q", expected, sum)
	}
}
