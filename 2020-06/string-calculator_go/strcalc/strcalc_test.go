package strcalc_test

import (
	"errors"
	"testing"

	"github.com/GoFightNguyen/katas/2020-06/string-calculator_go/strcalc"
)

func TestAddReturns0WhenEmpty(t *testing.T) {
	expect := 0
	result, _ := strcalc.Add("")
	if expect != result {
		t.Errorf("Expected: %d, Actual: %d", expect, result)
	}
}

func TestAddReturnsItselfWhenSingleValue(t *testing.T) {
	expect := 1
	result, _ := strcalc.Add("1")
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
		result, _ := strcalc.Add(s.input)
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
		result, _ := strcalc.Add(s.input)
		if s.expect != result {
			t.Errorf("Input: %v, Expected: %d, Actual: %d", s.input, s.expect, result)
		}
	}
}

func TestAddErrorsForNegatives(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
		err    error
	}{
		{input: "-1", expect: "negatives not allowed: -1", err: errors.New("negatives not allowed: -1")},
		{input: "1,-2", expect: "negatives not allowed: -2", err: errors.New("negatives not allowed: -2")},
		{input: "-1,2,-3", expect: "negatives not allowed: -1,-3", err: errors.New("negatives not allowed: -1,-3")},
	}

	for _, s := range scenarios {
		_, err := strcalc.Add(s.input)
		if err.Error() != s.err.Error() {
			t.Errorf("Input: %v, Expected: %v, Actual: %v", s.input, s.err, err)
		}
	}
}

func TestAddIgnoresNumbersGreaterThan1000(t *testing.T) {
	scenarios := []struct {
		input  string
		expect int
	}{
		{input: "1,1000", expect: 1001},
		{input: "1000,1", expect: 1001},
		{input: "1,1001", expect: 1},
		{input: "1001,1", expect: 1},
		{input: "1,1000,2,1001", expect: 1003},
		{input: "1001", expect: 0},
	}

	for _, s := range scenarios {
		result, _ := strcalc.Add(s.input)
		if s.expect != result {
			t.Errorf("Input: %v, Expected: %d, Actual: %d", s.input, s.expect, result)
		}
	}
}
