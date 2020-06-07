package strcalc

import (
	"strconv"
	"strings"
)

// Add returns the summation of the numbers provided.
func Add(numbers string) int {
	result := 0

	numbers = strings.ReplaceAll(numbers, "\n", ",")

	for _, s := range strings.Split(numbers, ",") {
		n, _ := strconv.Atoi(s)
		result += n
	}

	return result
}
