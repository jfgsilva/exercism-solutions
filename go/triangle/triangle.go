// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	NaT = "not a triangle" // not a triangle
	Equ = "equilateral"    // equilateral
	Iso = "isosceles"      // isosceles
	Sca = "scalene"        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case a+b+c == 0:
		return NaT
	case a <= 0 || b <= 0 || c <= 0:
		return NaT
	case a+b < c || a+c < b || b+c < a || c+a < b:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	default:
		return Sca
	}
}
