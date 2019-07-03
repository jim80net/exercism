// Package triangle determines the type of triangle
package triangle

import "math"

// Kind is one of the constants below
type Kind string

// Pick values for the following identifiers used by the test program.
const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// NumberOfSidesEqual compares each side with the others, incrementing count if theres a match
func NumberOfSidesEqual(a, b, c float64) int {
	count := 0
	if a == b || a == c {
		count++
	}
	if b == a || b == c {
		count++
	}
	if c == a || c == b {
		count++
	}
	return count
}

// PositiveLengthTest ensures triangle has positive length sides
func PositiveLengthTest(a, b, c float64) bool {
	if a <= 0 {
		return false
	}
	if b <= 0 {
		return false
	}
	if c <= 0 {
		return false
	}
	return true
}

// InequalityTest ensures sum of lengths of any two sides are >= third side
func InequalityTest(a, b, c float64) bool {
	if a+b < c {
		return false
	}
	if a+c < b {
		return false
	}
	if b+c < a {
		return false
	}
	return true
}

//IsNaN ensures all three sides are numbers
func IsNaN(a, b, c float64) bool {
	if math.IsNaN(a) {
		return false
	}
	if math.IsNaN(b) {
		return false
	}
	if math.IsNaN(c) {
		return false
	}
	return true
}

//IsInf ensures all three sides are numbers
func IsInf(a, b, c float64) bool {
	if math.IsInf(a, 0) {
		return false
	}
	if math.IsInf(b, 0) {
		return false
	}
	if math.IsInf(c, 0) {
		return false
	}
	return true
}

// KindFromSides determines the type of triangle
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if NumberOfSidesEqual(a, b, c) == 3 {
		k = Equ
	}
	if NumberOfSidesEqual(a, b, c) == 2 {
		k = Iso
	}
	if NumberOfSidesEqual(a, b, c) == 0 {
		k = Sca
	}
	if !PositiveLengthTest(a, b, c) {
		k = NaT
	}
	if !InequalityTest(a, b, c) {
		k = NaT
	}
	if !IsNaN(a, b, c) {
		k = NaT
	}
	if !IsInf(a, b, c) {
		k = NaT
	}

	return k
}
