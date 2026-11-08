// Package leap comprises a function that tells whether a year is a leap year
package leap

// IsLeapYear determines whether a given year is a leap year.
func IsLeapYear(year int) bool {

	var check bool

	if year%4 == 0 {
		check = true
	}

	if year%100 == 0 {
		check = false
	}

	if year%400 == 0 {
		check = true
	}

	return check
}
