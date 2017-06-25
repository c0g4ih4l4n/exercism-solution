package triangle

import (
	"math"
)

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 1) ||
		math.IsInf(b, 1) || math.IsInf(c, 1) || math.IsInf(a, -1) ||
		math.IsInf(b, -1) || math.IsInf(c, -1) || a <= 0 || b <= 0 || c <= 0 ||
		a+b < c || b+c < a || a+c < b {
		return NaT
	}

	if (a-b)*(b-c)*(c-a) != 0 {
		return Sca
	} else if a == b && b == c {
		return Equ
	} else {
		return Iso
	}
}

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

// Pick values for the following identifiers used by the test program.
const NaT = 1 // not a triangle
const Equ = 2 // equilateral
const Iso = 3 // isosceles
const Sca = 4 // scalene
