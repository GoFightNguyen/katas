package soundex

import "testing"

func TestRobertReturnsR163(t *testing.T) {
	actual := Index("Robert")
	expect := "R163"
	if actual != expect {
		t.Errorf("Expected: %v, Actual: %v", expect, actual)
	}
}

func TestRupertReturnsR163(t *testing.T) {
	actual := Index("Rupert")
	expect := "R163"
	if actual != expect {
		t.Errorf("Expected: %v, Actual: %v", expect, actual)
	}
}

func TestRubinReturnsR150(t *testing.T) {
	actual := Index("Rubin")
	expect := "R150"
	if actual != expect {
		t.Errorf("Expected: %v, Actual: %v", expect, actual)
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
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "abcdefghijklmnopqrstuvwxyz", expect: "a123e12hi22455o12623u1w2y2"},
		{input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", expect: "A123E12HI22455O12623U1W2Y2"},
	}

	for _, s := range scenarios {
		actual := encodeLabialConsonants(s.input)
		expect := s.expect
		if actual != expect {
			t.Errorf("Expected: %v, Actual: %v", expect, actual)
		}
	}
}

func TestRemoveVowelLikeLetters(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "abcdefghijklmnopqrstuvwxyz", expect: "bcdfgjklmnpqrstvxz"},
		{input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", expect: "BCDFGJKLMNPQRSTVXZ"},
	}

	for _, s := range scenarios {
		actual := removeVowelLikeLetters(s.input)
		expect := s.expect
		if actual != expect {
			t.Errorf("Expected: %v, Actual: %v", expect, actual)
		}
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
