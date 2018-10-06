package triangle

import (
	"math"
)

//Kind is a string
type Kind string

//these are the constatnt used in this programe
const (
	NaT = "NaT"
	Equ = "Equ"
	Iso = "Iso"
	Sca = "Sca"
)

// KindFromSides returns traingle type
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if a <= 0 || b <= 0 || c <= 0 || math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) || (a+b < c) || (a+c < b) || (b+c < a) {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || b == c || a == c {
		k = Iso
	} else if a != b || b != c || c != a {
		k = Sca
	}
	return k
}
