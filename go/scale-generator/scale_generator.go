package scale

import (
	"strings"
)

// Define groups array
var useSharp = map[string]string{
	"G":  "G",
	"D":  "D",
	"A":  "A",
	"E":  "E",
	"B":  "B",
	"F#": "F#",
	"e":  "e",
	"b":  "b",
	"f#": "f#",
	"c#": "c#",
	"g#": "g#",
	"d#": "d#",
	"a":  "a",
	"C":  "C",
}

// var useFlats = map[string]string{
// 	"F":  "F",
// 	"Bb": "Bb",
// 	"Eb": "Eb",
// 	"Ab": "Ab",
// 	"Db": "Db",
// 	"Gb": "Gb",
// 	"g":  "g",
// 	"c":  "c",
// 	"f":  "f",
// 	"bb": "bb",
// 	"eb": "eb",
// 	"d":  "d",
// }

func NextMTone(index int, chromatic []string) int {
	if len(chromatic) == index {
		index = 0
	}

	count := 0

	for i := index; ; i++ {
		if i == len(chromatic) {
			i = 0
		}

		if count == 1 {
			return i
		}

		count++
	}
}

func NextmTone(index int, chromatic []string) int {
	if len(chromatic) == index {
		index = 0
	}

	return index
}

func NextATone(index int, chromatic []string) int {
	if len(chromatic) == index {
		index = 0
	}

	count := 0

	for i := index; ; i++ {
		if i == len(chromatic) {
			i = 0
		}

		if count == 2 {
			return i
		}

		count++
	}
}

func Scale(tonic string, interval string) []string {
	var ret []string
	var defaultChromatic = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	var chromaticMajor = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

	_, ok := useSharp[tonic]

	if !ok {
		defaultChromatic = chromaticMajor
	}

	// locate index of tonic
	var index int

	for i, elem := range defaultChromatic {
		if strings.ToLower(elem) == strings.ToLower(tonic) {
			index = i
			break
		}
	}

	ret = append(ret, defaultChromatic[index])

	index++

	if index == len(defaultChromatic) {
		index = 0
	}

	intervalIndex := 0
	hasInterval := len(interval) > 0

	for i := index; ; i++ {
		if i == len(defaultChromatic) {
			i = 0
		}

		if strings.ToLower(defaultChromatic[i]) == strings.ToLower(tonic) {
			break
		}

		if intervalIndex == len(interval)-1 {
			break
		}

		if hasInterval {
			var nextToneIndex int

			switch string(interval[intervalIndex]) {
			case "M":
				nextToneIndex = NextMTone(i, defaultChromatic)

			case "m":
				nextToneIndex = NextmTone(i, defaultChromatic)

			case "A":
				nextToneIndex = NextATone(i, defaultChromatic)
			}

			i = nextToneIndex
			intervalIndex++
		}

		ret = append(ret, defaultChromatic[i])
	}

	return ret
}
