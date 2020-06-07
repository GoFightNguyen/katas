package strcalc

import (
	"errors"
	"strconv"
	"strings"
)

// Add returns the summation of the numbers provided.
// Numbers greater than 1000 are ignored.
// Negatives will cause an error.
func Add(numbers string) (sum int, err error) {
	delimiterLine, numbers := separateDelimiterLineFromNumbers(numbers)
	delimiter := parseDelimiter(delimiterLine)
	numbers = strings.ReplaceAll(numbers, delimiter, ",")
	splitNumbers := strings.Split(numbers, ",")

	err = validateNoNegatives(splitNumbers)
	if err != nil {
		return
	}

	sum = sumNumbersLessThan1001(splitNumbers)

	return
}

func validateNoNegatives(numbers []string) error {
	negatives := getNegatives(numbers)

	if len(negatives) > 0 {
		msg := "negatives not allowed: " + strings.Join(negatives, ",")
		return errors.New(msg)
	}

	return nil
}

func getNegatives(numbers []string) []string {
	var negatives []string
	for _, n := range numbers {
		if strings.HasPrefix(n, "-") {
			negatives = append(negatives, n)
		}
	}

	return negatives
}

func sumNumbersLessThan1001(numbers []string) int {
	sum := 0
	for _, s := range numbers {
		n, _ := strconv.Atoi(s)
		if n > 1000 {
			continue
		}

		sum += n
	}

	return sum
}

func separateDelimiterLineFromNumbers(numbers string) (delimiterLine string, numbersLine string) {
	if strings.HasPrefix(numbers, "//") {
		split := strings.Split(numbers, "\n")
		delimiterLine = split[0]
		numbersLine = split[1]
	} else {
		numbersLine = numbers
	}

	return
}

func parseDelimiter(delimiterLine string) (delimiter string) {
	if delimiterLine != "" {
		if strings.Contains(delimiterLine, "[") {
			endIndex := strings.Index(delimiterLine, "]")
			delimiter = delimiterLine[3:endIndex]
		} else {
			delimiter = delimiterLine[2:]
		}
	} else {
		delimiter = "\n"
	}

	return
}
