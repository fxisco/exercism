package luhn

import (
	"regexp"
)

func Valid(str string) (result bool) {

	if len(str) > 0 {
		// Remover los caracteres que no sean numericos
		rule := regexp.MustCompile(`[^0-9]`)
		newStr := rule.ReplaceAllString(str, "")

		var total int
		strLen := len(newStr) - 1

		for i := strLen; i >= 0; i-- {
			digit := int(newStr[i] - '0')

			// Double every second digit, starting from the right
			if i < strLen && (strLen-i)%2 == 1 {
				var newDigit int
				doubleDigit := digit * 2

				// Check if adding is greater than 9, if so substract 9 to the resutl
				if doubleDigit > 9 {
					newDigit = doubleDigit - 9
				} else {
					newDigit = doubleDigit
				}

				digit = newDigit
			}

			total += digit
		}

		// Check if the sum of all digits % 10 is equal 0

		if total > 0 && (total%10) == 0 {
			result = true
		}
	}

	return
}
