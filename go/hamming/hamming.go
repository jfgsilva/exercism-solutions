package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Distances don't match")
	}
	mismatches := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			mismatches += 1
		}
	}
	return mismatches, nil
}
