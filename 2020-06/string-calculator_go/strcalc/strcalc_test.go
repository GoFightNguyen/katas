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

func TestAddReturnsSumOfCommaDelimitedNumbers(t *testing.T) {
	scenarios := []struct {
		input  string
		expect int
	}{
		{input: "1,2", expect: 3},
		{input: "1,2,4,6", expect: 13},
	}

	for _, s := range scenarios {
		result := strcalc.Add(s.input)
		if s.expect != result {
			t.Errorf("Scenario: %v, Expected: %d, Actual: %d", s.input, s.expect, result)
		}
	}
}
