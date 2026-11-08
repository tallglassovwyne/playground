package hamming

import "errors"

// compares two strands of DNA and counts the differences between them.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Hamming distance is only defined for sequences of equal length")
	}

	var distance int = 0

	for i, j := 0, 0; i < len(a); i, j = i+1, j+1 {
		if a[i] != b[j] {
			distance += 1
		}
	}

	return distance, nil
}
