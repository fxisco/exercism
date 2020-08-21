package raindrops

import "fmt"

var factors = []struct {
	factor  int
	message string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(input int) (result string) {
	var condition bool

	for _, f := range factors {
		if input%f.factor == 0 {
			result += f.message
			condition = true
		}
	}

	if !condition {
		return fmt.Sprintf("%d", input)
	}

	return
}
