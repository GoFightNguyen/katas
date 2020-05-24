package soundex

import "strings"

func Encode(name string) string {
	firstLetter := saveFirstLetter(name)
	name = name[1:]
	name = encodeVowelSounds(name)
	name = encodeConsonants(name)
	name = replaceRepeatingAdjacentDigits(name)
	name = removeAllZeroDigits(name)
	name = rightPadUntilThreeDigits(name)
	name = trimToThreeDigits(name)
	return firstLetter + name
}

func saveFirstLetter(name string) string {
	return name[:1]
}

func encodeVowelSounds(name string) string {
	vowelSounds := [16]string{
		"a", "e", "i", "o", "u", "y", "h", "w",
		"A", "E", "I", "O", "U", "Y", "H", "W",
	}

	for _, letter := range vowelSounds {
		name = strings.ReplaceAll(name, letter, "0")
	}

	return name
}

func encodeConsonants(name string) string {
	consonants := map[string]string{
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

	for letter, digit := range consonants {
		name = strings.ReplaceAll(name, letter, digit)
	}

	return name
}

func replaceRepeatingAdjacentDigits(name string) string {
	digitsToKeep := []rune{}
	for _, digit := range name {

		if len(digitsToKeep) == 0 {
			digitsToKeep = append(digitsToKeep, digit)
			continue
		}

		if digitsToKeep[len(digitsToKeep)-1] != digit {
			digitsToKeep = append(digitsToKeep, digit)
		}
	}

	return string(digitsToKeep)
}

func removeAllZeroDigits(name string) string {
	return strings.ReplaceAll(name, "0", "")
}

func rightPadUntilThreeDigits(name string) string {
	for len(name) < 3 {
		name += "0"
	}

	return name
}

func trimToThreeDigits(name string) string {
	if len(name) >= 3 {
		return name[0:3]
	}

	return name[0:]
}
