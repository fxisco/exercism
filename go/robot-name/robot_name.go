package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
	RobotName string
}

const (
	MIN = 65
	MAX = 90
)

var assigned = map[string]int{}

func (r *Robot) Name() (string, error) {
	if len(r.RobotName) > 0 {
		return r.RobotName, nil
	}

	var result string
	finish := false

	for !finish {
		result = ""

		for i := 0; i < 2; i++ {
			first := rand.Intn(MAX-MIN) + MIN
			result += string(rune(first))
		}

		item, _ := assigned[result]
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
