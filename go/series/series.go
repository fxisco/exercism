package series

func All(n int, s string) []string {
	result := make([]string, 0)

	for i := 0; i <= len(s)-n; i++ {
		serie := s[i : i+n]

		result = append(result, serie)
	}

	return result
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}
