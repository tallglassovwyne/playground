// Package luhn implements the Luhn algorithm on an identification number.
// The Luhn algorithm, or Luhn formula (creator: IBM scientist Hans Peter Luhn),
// also known as the "modulus 10" or "mod 10" algorithm, is a simple check digit formula used to validate a variety of identification numbers.
// The purpose is to design a numbering scheme in such a way that when a human is entering a number, a computer can quickly check it for errors.
package luhn

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

// Valid determines whether a number is valid according to the Luhn formula.
func Valid(id string) bool {

    // Trim whitespace from the string.
	id = strings.ReplaceAll(id, " ", "")

	// Strings of length 1 or less are not valid.
	if len(id) <= 1 {
		return false
	}

	// Checks if id contains any characters besides digits.
	for _, r := range id {
		if r < '0' || r > '9' {
			return false
		}
	}

	var arrayOfNumbers []string
	var integersArray []int

	for _, r := range id {
		arrayOfNumbers = append(arrayOfNumbers, string(r))
	}

	for i := 0; i < len(arrayOfNumbers); i++ {
		num, err := strconv.Atoi(arrayOfNumbers[i])

		if err != nil {
			log.Fatal(err)
		}

		integersArray = append(integersArray, num)
	}

	slices.Reverse(integersArray)

	for i := 0; i < len(integersArray); i++ {
		if i % 2 != 0 {
			if (integersArray[i] * 2) > 9 {
				integersArray[i] = integersArray[i] * 2 - 9
			} else {
				integersArray[i] = integersArray[i] * 2
			}
		} 
	}

	total := 0

	for i := 0; i < len(integersArray); i++ {
		total += integersArray[i]
	}

	if total % 10 == 0 {
		return true
	}

	return false
}