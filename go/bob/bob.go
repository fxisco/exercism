// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

func Hey(remark string) string {

	// Cleanse string
	cleansedRemark := strings.TrimSpace(remark)

	// if empty -> Fine. Be that way!
	if len(cleansedRemark) == 0 {
		return "Fine. Be that way!"
	}

	allUppercaseRe := regexp.MustCompile(`[^A-Z]`)
	allLettersRe := regexp.MustCompile(`[^A-Za-z]`)
	str := regexp.MustCompile(`\s+`)

	cleansedAllUppercase := allUppercaseRe.ReplaceAllString(cleansedRemark, "")
	cleansedAllLetters := allLettersRe.ReplaceAllString(cleansedRemark, "")
	testAllUppercase := strings.TrimSpace(str.ReplaceAllString(cleansedAllUppercase, " "))
	testAllLetters := strings.TrimSpace(str.ReplaceAllString(cleansedAllLetters, " "))

	if len(testAllUppercase) == len(testAllLetters) && len(testAllUppercase) > 0 {

		// if yelling a question
		if strings.HasSuffix(cleansedRemark, "?") {
			return "Calm down, I know what I'm doing!"
		}

		return "Whoa, chill out!"
	}

	// any question(?) -> Sure.
	if strings.HasSuffix(cleansedRemark, "?") {
		return "Sure."
	}

	// default -> Whatever.
	return "Whatever."
}
