package accumulate

func Accumulate(input []string, converter func(string) string) []string {
	var slice = make([]string, len(input))

	for index, text := range input {
		slice[index] = converter(text)
	}

	return slice
}
