// Package isbnverifier validates book identification numbers.
package isbnverifier

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

// Helper function that whether a character is a member of an array.
func contains(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

// IsValidISBN returns whether a book's identification number is valid.
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	characters := []string {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	checkDigit := []string {"X", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	if len(isbn) > 10 || len(isbn) < 10 {
		return false
	}

	if !slices.Contains(checkDigit, string(isbn[len(isbn) - 1])) {
		return false
	}

	for _, char := range isbn[0:len(isbn)-1] {
		charStr := string(char)

		if contains(characters, charStr) {
			continue
		} else {
			return false
		}
	}

	index, total := 10, 0

	for _, char := range isbn[:len(isbn) - 1] {
		num, err := strconv.Atoi(string(char))

		if err != nil {
			log.Fatal(err)
		}

		total += num * index

		index -= 1
	}

	if isbn[len(isbn) -1] == 'X' {
		total += 10 * 1
	} else {
		num, err := strconv.Atoi(string(isbn[len(isbn) -1]))

		if err != nil {
			log.Fatal(err)
		}

		total += num * 1
	}

	if total%11 == 0{
		return true
	}

	return false
}