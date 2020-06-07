package strcalc

import (
	"strconv"
	"strings"
)

// Add returns the summation of the numbers provided.
func Add(numbers string) int {
	result := 0

	if strings.HasPrefix(numbers, "//") {
		split := strings.Split(numbers, "\n")
		delimiter := split[0][2:]
		numbers = strings.ReplaceAll(split[1], delimiter, ",")
	} else {
		numbers = strings.ReplaceAll(numbers, "\n", ",")
	}

	for _, s := range strings.Split(numbers, ",") {
		n, _ := strconv.Atoi(s)
		result += n
	}

	return result
}
