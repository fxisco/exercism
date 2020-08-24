package reverse

import (
	"strings"
)

func Reverse(text string) (result string) {
	aux := strings.Split(text, "")

	for i := len(aux) - 1; i >= 0; i-- {
		result += aux[i]
	}

	return
}
