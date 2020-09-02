package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(arr []string) FreqMap {
	switch len(arr) {
	case 0:
		return FreqMap{}
	case 1:
		return Frequency(arr[0])
	}

	ch := make(chan FreqMap)

	f := func(l []string) {
		ch <- ConcurrentFrequency(l)
	}

	half := len(arr) / 2

	go f(arr[:half])
	go f(arr[half:])

	m := <-ch

	for r, n := range <-ch {
		m[r] += n
	}

	return m
}
