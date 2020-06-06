package strcalc

import (
	"strconv"
	"strings"
)

// Add returns the summation of the numbers provided.
func Add(numbers string) int {
	if numbers == "" {
		return 0
	}

	result := 0
	for _, s := range strings.Split(numbers, ",") {
		n, _ := strconv.Atoi(s)
		result += n
	}

	return result
}
