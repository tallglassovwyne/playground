// Package collatz_conjecture counts how many steps a number needs to reach 1.
// Under the Collatz rules: halve if it's even, otherwise triple it and add 1.
package collatz

import "errors"

// CollatzConjecture returns the number of steps n takes to reach 1
// For example, 1 takes 0 steps, and 12 takes 9 steps
// It returns an error if n is less than 1.
func Steps(n int) (int, error) {
	numberOfSteps := 0

	// The conjecture is only defined for positive integers. Zero and negative
	// values never reach 1, so reject them rather than loop forever.
	if n < 1 {
		return 0, errors.New("only positive integers can be tested")
	}

	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = (3 * n) + 1

		}

		numberOfSteps++

	}

	return numberOfSteps, nil
}
