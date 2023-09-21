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

func Test5(t *testing.T) {
	sum := Calculator("10,20,30")
	expected := 60
	assertEqual(t, sum, expected)
}

func Test6(t *testing.T) {
	sum := Calculator("1\n2")
	expected := 3
	assertEqual(t, sum, expected)
}

func assertEqual(t *testing.T, sum int, expected int) {
	if sum != expected {
		t.Errorf("Expected %v but got %v", expected, sum)
	}
}
