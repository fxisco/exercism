package etl

import (
	"strings"
)

func Transform(given map[int][]string) map[string]int {
	result := make(map[string]int)

	for key, elements := range given {
		for _, letter := range elements {
			result[strings.ToLower(letter)] = key
		}
	}

	return result
}
