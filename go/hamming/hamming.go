package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("can't work if they are not the same size")
	}

	s := strings.Split(a, "")
	var counter int

	for i, char := range s {
		if string(char) != string(b[i]) {
			counter++
		}
	}

	return counter, nil
}
