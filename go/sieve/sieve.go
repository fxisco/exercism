package sieve

func Sieve(num int) []int {
	result := make([]int, 0)

	if num < 2 {
		return result
	}

	if num == 2 {
		return []int{2}
	}

	nSlots := num - 1
	availables := make([]bool, nSlots)
	i := 2
	mul := 1

	for i < num {
		index := (i * mul) - 2

		if index >= nSlots {
			i++
			mul = 1
			continue
		}

		if availables[index] == true {
			mul++
			continue
		}

		if mul > 1 {
			availables[index] = true
		}

		mul++
	}

	for index, element := range availables {
		if element == false {
			result = append(result, index+2)
		}
	}

	return result
}
