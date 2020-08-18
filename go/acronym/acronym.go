// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	hyphenRe := regexp.MustCompile(`-`)
	allLettersRe := regexp.MustCompile(`[^A-Za-z ]`)
	str := regexp.MustCompile(`\s+`)

	cleansedText := hyphenRe.ReplaceAllString(s, " ")
	cleansedText = allLettersRe.ReplaceAllString(cleansedText, "")
	allLetters := strings.TrimSpace(str.ReplaceAllString(cleansedText, " "))
	arr := strings.Split(strings.ToUpper(allLetters), " ")
	result := ""

	for _, word := range arr {
		result += word[0:1]
	}

	return result
}
