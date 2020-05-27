package soundex

import "strings"

func Index(name string) string {
	firstLetter := getTheFirstLetter(name)
	encoding := name[1:]
	encoding = encodeLabialConsonants(encoding)
	encoding = removeVowelLikeLetters(encoding)
	encoding = rightPadToThreeDigits(encoding)
	return firstLetter + encoding
}

func getTheFirstLetter(name string) string {
	return name[0:1]
}

func encodeLabialConsonants(name string) string {
	mapping := map[string]string{
		"b": "1",
		"f": "1",
		"p": "1",
		"v": "1",
		"c": "2",
		"g": "2",
		"j": "2",
		"k": "2",
		"q": "2",
		"s": "2",
		"x": "2",
		"z": "2",
		"d": "3",
		"t": "3",
		"l": "4",
		"m": "5",
		"n": "5",
		"r": "6",

		"B": "1",
		"F": "1",
		"P": "1",
		"V": "1",
		"C": "2",
		"G": "2",
		"J": "2",
		"K": "2",
		"Q": "2",
		"S": "2",
		"X": "2",
		"Z": "2",
		"D": "3",
		"T": "3",
		"L": "4",
		"M": "5",
		"N": "5",
		"R": "6",
	}

	for key, encoding := range mapping {
		name = strings.ReplaceAll(name, key, encoding)
	}

	return name
}

func removeVowelLikeLetters(name string) string {
	vowelLikeLetters := [16]string{
		"a", "e", "i", "o", "u", "y", "h", "w",
		"A", "E", "I", "O", "U", "Y", "H", "W",
	}

	for _, letter := range vowelLikeLetters {
		name = strings.ReplaceAll(name, letter, "")
	}

	return name
}

func rightPadToThreeDigits(name string) string {
	for len(name) < 3 {
		name = name + "0"
	}
	return name
}
