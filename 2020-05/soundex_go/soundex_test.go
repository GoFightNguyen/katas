package soundex

import "testing"

func TestRobertReturnsR163(t *testing.T) {
	scenarios := [3]string{
		"Robert",
		"ROBERT",
		"robert",
	}
	for _, name := range scenarios {
		actual := Index(name)
		expect := "R163"
		if actual != expect {
			t.Errorf("Name: %v, Expected: %v, Actual: %v", name, expect, actual)
		}
	}
}

func TestRupertReturnsR163(t *testing.T) {
	scenarios := [3]string{
		"Rupert",
		"RUPERT",
		"rupert",
	}
	for _, name := range scenarios {
		actual := Index(name)
		expect := "R163"
		if actual != expect {
			t.Errorf("Name: %v, Expected: %v, Actual: %v", name, expect, actual)
		}
	}
}

func TestRubinReturnsR150(t *testing.T) {
	scenarios := [3]string{
		"Rubin",
		"RUBIN",
		"rubin",
	}
	for _, name := range scenarios {
		actual := Index(name)
		expect := "R150"
		if actual != expect {
			t.Errorf("Name: %v, Expected: %v, Actual: %v", name, expect, actual)
		}
	}
}

func TestAshcraftReturnsA261(t *testing.T) {
	scenarios := [3]string{
		"Ashcraft",
		"ASHCRAFT",
		"ashcraft",
	}
	for _, name := range scenarios {
		actual := Index(name)
		expect := "A261"
		if actual != expect {
			t.Errorf("Name: %v, Expected: %v, Actual: %v", name, expect, actual)
		}
	}
}

func TestAshcroftReturnsA261(t *testing.T) {
	scenarios := [3]string{
		"Ashcroft",
		"ASHCROFT",
		"ashcroft",
	}
	for _, name := range scenarios {
		actual := Index(name)
		expect := "A261"
		if actual != expect {
			t.Errorf("Name: %v, Expected: %v, Actual: %v", name, expect, actual)
		}
	}
}

func TestGetTheFirstLetter(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Robert", expect: "R"},
		{input: "Allen", expect: "A"},
	}

	for _, s := range scenarios {
		actual := getTheFirstLetter(s.input)
		expect := s.expect
		if actual != expect {
			t.Errorf("Expected: %v, Actual: %v", expect, actual)
		}
	}
}

func TestEncodeLabialConsonants(t *testing.T) {
	actual := encodeLabialConsonants("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	expect := "A123E12HI22455O12623U1W2Y2"
	if actual != expect {
		t.Errorf("Expected: %v, Actual: %v", expect, actual)
	}
}

func TestRemoveVowelLikeLetters(t *testing.T) {

	actual := removeVowelLikeLetters("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	expect := "BCDFGJKLMNPQRSTVXZ"
	if actual != expect {
		t.Errorf("Expected: %v, Actual: %v", expect, actual)
	}
}

func TestRightPadToThreeDigits(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "1", expect: "100"},
		{input: "12", expect: "120"},
		{input: "123", expect: "123"},
		{input: "1234", expect: "1234"},
	}

	for _, s := range scenarios {
		actual := rightPadToThreeDigits(s.input)
		expect := s.expect
		if actual != expect {
			t.Errorf("Expected: %v, Actual: %v", expect, actual)
		}
	}
}

func TestReduceLabialConsonantsSeparatedByHToSingleDigit(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "1", expect: "1"},
		{input: "H", expect: "H"},
		{input: "H1", expect: "H1"},
		{input: "1H", expect: "1H"},
		{input: "1H2", expect: "1H2"},
		{input: "1H2H1", expect: "1H2H1"},
		{input: "1H1", expect: "1"},
		{input: "1H1H1", expect: "1"},
		{input: "1H2H2", expect: "1H2"},
	}

	for _, s := range scenarios {
		actual := reduceLabialConsonantsSeparatedByHToSingleDigit(s.input)
		expect := s.expect
		if actual != expect {
			t.Errorf("Expected: %v, Actual: %v", expect, actual)
		}
	}
}

func TestTrimToThreeDigits(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "1", expect: "1"},
		{input: "123", expect: "123"},
		{input: "1234", expect: "123"},
		{input: "12345", expect: "123"},
	}

	for _, s := range scenarios {
		actual := trimToThreeDigits(s.input)
		expect := s.expect
		if actual != expect {
			t.Errorf("Expected: %v, Actual: %v", expect, actual)
		}
	}
}
