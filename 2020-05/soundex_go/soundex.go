package soundex

import "strings"

func Index(name string) string {
	name = strings.ToUpper(name)

	firstLetter := getTheFirstLetter(name)
	encoding := name[1:]
	encoding = encodeLabialConsonants(encoding)
	encoding = reduceLabialConsonantsSeparatedByHToSingleDigit(encoding)
	encoding = removeVowelLikeLetters(encoding)
	encoding = rightPadToThreeDigits(encoding)
	encoding = trimToThreeDigits(encoding)
	return firstLetter + encoding
}

func getTheFirstLetter(name string) string {
	return name[0:1]
}

func encodeLabialConsonants(name string) string {
	mapping := map[string]string{
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

func trimToThreeDigits(name string) string {
	if len(name) > 3 {
		return name[0:3]
	}

	return name
}

func reduceLabialConsonantsSeparatedByHToSingleDigit(name string) string {
	var previous rune
	replacements := []string{}

	for _, current := range name {
		if current == 'H' {
			continue
		}

		if previous == current {
			replacements = append(replacements, string(current))
		}

		previous = current
	}

	for _, digit := range replacements {
		name = strings.ReplaceAll(name, digit+"H"+digit, digit)
	}

	if len(replacements) > 0 {
		name = reduceLabialConsonantsSeparatedByHToSingleDigit(name)
	}

	return name
}
