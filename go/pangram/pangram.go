package pangram

import (
	"regexp"
	"strings"
)

func IsPangram(s string) (response bool) {
	s = strings.ToLower(s)
	alphabetRe := regexp.MustCompile(`[^a-z]`)

	cleansedText := alphabetRe.ReplaceAllString(s, "")
	dic := make(map[string]string)

	for _, letter := range cleansedText {
		dic[string(letter)] = string(letter)
	}

	if len(dic) != 26 {
		return
	}

	return true
}
