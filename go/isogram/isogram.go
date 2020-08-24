package isogram

import (
	"strings"
)

func IsIsogram(str string) bool {
	repetitions := make(map[string]int)
	letters := strings.Split(strings.ToLower(str), "")
	result := true

	for _, l := range letters {
		_, isPresent := repetitions[l]

		if isPresent && l != "-" && l != " " {
			result = false
			break
		} else {
			repetitions[l] = 1
		}
	}

	return result
}
