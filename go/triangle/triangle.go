// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "Nat" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Condition 1: all sizes lenght > 0
	// L1 + L2 >= L3
	sides := []float64{a, b, c}

	for _, element := range sides {
		if math.IsNaN(element) || math.IsInf(element, 1) || math.IsInf(element, -1) || element <= 0 {
			return NaT
		}
	}

	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	sidesMap := make(map[float64]float64)

	sidesMap[a] = a
	sidesMap[b] = b
	sidesMap[c] = c

	if len(sidesMap) == 2 {
		return Iso
	}

	if len(sidesMap) == 1 {
		return Equ
	}

	return Sca
}
