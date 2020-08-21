package scrabble

import (
	"regexp"
	"strings"
)

var letterValues = []struct {
	letters string
	value   int
}{
	{
		`A|E|I|O|U|L|N|R|S|T`,
		1,
	},
	{
		`D|G`,
		2,
	},
	{
		`B|C|M|P`,
		3,
	},
	{
		`F|H|V|W|Y`,
		4,
	},
	{
		`K`,
		5,
	},
	{
		`J|X`,
		8,
	},
	{
		`Q|Z`,
		10,
	},
}

func Score(text string) (result int) {
	for _, c := range letterValues {
		rule := regexp.MustCompile(c.letters)
		matches := rule.FindAllStringIndex(strings.ToUpper(text), -1)

		result += len(matches) * c.value
	}

	return
}
