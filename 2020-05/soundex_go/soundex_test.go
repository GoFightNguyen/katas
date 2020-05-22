package soundex

import "testing"

func TestRobertIsEncodedAsR163(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Robert", expect: "R163"},
		{input: "Rupert", expect: "R163"},
		{input: "Rubin", expect: "R150"},
	}

	for _, s := range scenarios {
		result := Encode(s.input)
		expect := s.expect
		if result != expect {
			t.Errorf("Expected: %q\nActual: %q", expect, result)
		}
	}
}

func TestSaveFirstLetterReturnsTheFirstLetterOfTheName(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Robert", expect: "R"},
		{input: "Allen", expect: "A"},
	}

	for _, s := range scenarios {
		result := saveFirstLetter(s.input)
		expect := s.expect
		if result != expect {
			t.Errorf("Expected: %q\nActual: %q", expect, result)
		}
	}
}

func TestEncodeVowelSoundsReplacesThemWithZero(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "abcdefghijklmnopqrstuvwxyz", expect: "0bcd0fg00jklmn0pqrst0v0x0z"},
		{input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", expect: "0BCD0FG00JKLMN0PQRST0V0X0Z"},
	}

	for _, s := range scenarios {
		result := encodeVowelSounds(s.input)
		expect := s.expect
		if result != expect {
			t.Errorf("Expected: %q\nActual: %q", expect, result)
		}
	}
}

func TestEncodeConsonantsReplacesThemWithCorrectNumber(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "abcdefghijklmnopqrstuvwxyz", expect: "a123e12hi22455o12623u1w2y2"},
		{input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ", expect: "A123E12HI22455O12623U1W2Y2"},
	}

	for _, s := range scenarios {
		result := encodeConsonants(s.input)
		expect := s.expect
		if result != expect {
			t.Errorf("Expected: %q\nActual: %q", expect, result)
		}
	}
}

func TestReplaceRepeatingAdjacentDigitsWithSingleDigit(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "111", expect: "1"},
		{input: "121", expect: "121"},
		{input: "1221", expect: "121"},
		{input: "122344", expect: "1234"},
		{input: "122334", expect: "1234"},
	}

	for _, s := range scenarios {
		result := replaceRepeatingAdjacentDigits(s.input)
		expect := s.expect
		if result != expect {
			t.Errorf("Expected: %q\nActual: %q", expect, result)
		}
	}
}

func TestRemoveAllZeroDigits(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "101", expect: "11"},
		{input: "010", expect: "1"},
		{input: "1002", expect: "12"},
		{input: "102003", expect: "123"},
	}

	for _, s := range scenarios {
		result := removeAllZeroDigits(s.input)
		expect := s.expect
		if result != expect {
			t.Errorf("Expected: %q\nActual: %q", expect, result)
		}
	}
}

func TestRightPadUntilThreeDigits(t *testing.T) {

}
