// Package triangle determines the type of triangle
package triangle

import "math"

// Kind is one of the constants below
type Kind string

// Pick values for the following identifiers used by the test program.
const (
	NaT Kind = "Not A Triangle" // not a triangle
	Equ      = "Equilateral"    // equilateral
	Iso      = "Isoosceles"     // isosceles
	Sca      = "Scalar"         // scalene
)

// KindFromSides determines the type of triangle
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a+b+c) || math.IsInf(a+b+c, 0) {
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && a == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
