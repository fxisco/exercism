package cryptosquare

import (
	"fmt"
	"regexp"
	"strings"
)

func findCandR(len int) (int, int) {
	c := 1
	r := 1

	for i := 1; ; i++ {
		if c*r >= len {
			return c, r
		}

		if i%2 == 0 {
			r++
		} else {
			c++
		}
	}
}

func Encode(pt string) (ct string) {

	// Normalize text
	normalizeRe := regexp.MustCompile(`[^a-z0-9]`)
	cleansedText := normalizeRe.ReplaceAllString(strings.ToLower(pt), "")
	strLen := len(cleansedText)

	if len(cleansedText) == 0 {
		return
	}

	c, r := findCandR(strLen)
	padTxt := fmt.Sprintf("%-*s", c*r, cleansedText)

	chunks := make([]string, 0)

	for i := 0; i < r; i++ {
		piece := padTxt[i*c : (i+1)*c]
		chunks = append(chunks, piece)
	}

	encChunks := make([]string, 0)

	for i := 0; i < c; i++ {
		buff := ""

		for _, element := range chunks {
			buff += element[i : i+1]
		}

		encChunks = append(encChunks, buff)
	}

	ct = strings.Join(encChunks, " ")

	return
}
