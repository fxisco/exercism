package grains

import (
	"errors"
	"math"
)

func Total() (total uint64) {
	for i := 0; i < 64; i++ {
		total += uint64(math.Pow(2, float64(i)))
	}

	return
}

func Square(input int) (uint64, error) {
	if input <= 0 {
		return 0, errors.New("Should be greater than 0")
	}

	if input > 64 {
		return 0, errors.New("Should NOT be greater than 64")
	}

	return uint64(math.Pow(2, float64(input-1))), nil
}
