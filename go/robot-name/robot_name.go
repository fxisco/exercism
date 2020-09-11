package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
	RobotName string
	seen      map[string]int
}

const (
	MIN = 65
	MAX = 90
)

func (r *Robot) Name() (string, error) {
	var result string

	finish := false

	for !finish {
		for i := 0; i < 2; i++ {
			first := rand.Intn(MAX-MIN) + MIN
			result += string(rune(first))
		}

		item, _ := r.seen[result]
		item++

		if item < 1000 {
			result = fmt.Sprintf("%s%03d", result, item)
			finish = true
		}
	}

	r.RobotName = result

	return result, nil
}

func (r *Robot) Reset() (string, error) {
	r.RobotName = ""

	newName, err := r.Name()

	r.RobotName = newName

	return r.RobotName, err
}
