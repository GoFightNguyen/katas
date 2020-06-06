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

func TestAddReturnsSumOfTwoCommaSeparatedNumbers(t *testing.T) {
	expect := 3
	result := strcalc.Add("1,2")
	if expect != result {
		t.Errorf("Expected: %d, Actual: %d", expect, result)
	}
}
