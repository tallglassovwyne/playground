package hamming

import "errors"

// compares two strands of DNA and counts the differences between them.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Hamming distance is only defined for sequences of equal length")
	}

	var differenceCount int = 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			differenceCount += 1
		}
	}

	return differenceCount, nil
}
