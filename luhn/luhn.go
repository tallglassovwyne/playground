// Package luhn implements the Luhn algorithm on an identification number.
// The Luhn algorithm, or Luhn formula also known as the "modulus 10" or "mod 10" algorithm, 
// is a simple check digit formula used to validate a variety of identification numbers.
// The purpose is to design a numbering scheme in such a way that when a human is entering a number, a computer can quickly check it for errors.
package luhn

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

// removeWhiteSpaces is a helper function that removes ALL whitespaces from a string.
// It uses a list of space characters to remove respective spaces.
func removeWhiteSpaces(s string, checkValues ...string) string {

	for _, value := range checkValues {
		s = strings.ReplaceAll(s, value, "")
	}

	return strings.TrimSpace(s)
}

// Valid determines whether a number is valid according to the Luhn formula.
func Valid(id string) bool {

	id = strings.TrimSpace(id) // removes beginning and trailing whitespaces
	id = removeWhiteSpaces(id, " ")
	id = removeWhiteSpaces(id, "\t")
	id = removeWhiteSpaces(id, "\n")

	// strings of length 1 or less are marked as invalid.
	if len(id) <= 1 {
		return false
	}

	// checks if the id contains any non-digit string character.
	for _, r := range id {
		if r < '0' || r > '9' {
			return false
		}
	}

	var sliceOfNumbers []string
	var integersArray []int

	for _, r := range id {
		// splits the id string into a slice.
		sliceOfNumbers = append(sliceOfNumbers, string(r))
	}

	for i := range sliceOfNumbers {
		num, err := strconv.Atoi(sliceOfNumbers[i])

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