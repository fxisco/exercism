package lsproduct

import (
	"errors"
	"strconv"
	"strings"
)

func getMultiplication(digits string) (int64, error) {
	var result int64 = 1

	for _, digit := range strings.Split(digits, "") {
		n, err := strconv.ParseInt(digit, 10, 64)

		if err != nil {
			return -1, errors.New("Failed conversion")
		}

		result *= n
	}

	return result, nil
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return -1, errors.New("span SHOULD be >")
	}

	if span > len(digits) {
		return -1, errors.New("span SHOULD NOT be greater than #'s of digits")
	}

	max := int64(-1)

	for i := 0; i <= len(digits)-span; i++ {
		cDigit := digits[i : i+span]
		result, err := getMultiplication(cDigit)

		if err != nil {
			return -1, err
		}

		if result > max {
			max = result
		}
	}

	return max, nil
}
