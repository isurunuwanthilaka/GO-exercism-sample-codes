//hamming for identifying DNA patterns
package hamming

import (
	"errors"
)

//Distance for calculating distance between two DNA strings
func Distance(a, b string) (int, error) {

	lenA := len(a)
	lenB := len(b)
	if lenA != lenB {
		return -1, errors.New("Lengths do not match each other")
	}

	var r int
	r = 0

	for i := range a {
		if a[i] != b[i] {
			r++
		}
	}

	return r, nil
}
