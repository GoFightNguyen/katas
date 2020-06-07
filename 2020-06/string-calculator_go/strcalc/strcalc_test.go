package strcalc_test

import (
	"testing"

	"github.com/GoFightNguyen/katas/2020-06/string-calculator_go/strcalc"
)

func TestAddReturns0WhenEmpty(t *testing.T) {
	expect := 0
	result := strcalc.Add("")
	if expect != result {
		t.Errorf("Expected: %d, Actual: %d", expect, result)
	}
}

func TestAddReturnsItselfWhenSingleValue(t *testing.T) {
	expect := 1
	result := strcalc.Add("1")
	if expect != result {
		t.Errorf("Expected: %d, Actual: %d", expect, result)
	}
}

func TestAddSupportsDefaultDelimitersOfCommaAndNewLine(t *testing.T) {
	scenarios := []struct {
		input  string
		expect int
	}{
		{input: "1,2", expect: 3},
		{input: "1,2,4,6", expect: 13},
		{input: "1\n2", expect: 3},
		{input: "1\n2\n4\n6", expect: 13},
		{input: "1\n2,3", expect: 6},
	}

	for _, s := range scenarios {
		result := strcalc.Add(s.input)
		if s.expect != result {
			t.Errorf("Input: %v, Expected: %d, Actual: %d", s.input, s.expect, result)
		}
	}
}

func TestAddSupportsSpecifiedDelimiter(t *testing.T) {
	scenarios := []struct {
		input  string
		expect int
	}{
		{input: "//;\n1", expect: 1},
		{input: "//;\n1;2", expect: 3},
	}

	for _, s := range scenarios {
		result := strcalc.Add(s.input)
		if s.expect != result {
			t.Errorf("Input: %v, Expected: %d, Actual: %d", s.input, s.expect, result)
		}
	}
}
