// package phonenumber provides functions that clean up NANP (North America Numbering Plan) phone numbers.
package phonenumber

import (
	"errors"
	"strconv"
	"strings"
)

// removeCharacter is a helper function that receives a delimiter, removes the delimiter
// and returns the updated string.
func removeCharacter(s, delimiter string) string {
	if strings.HasPrefix(s, delimiter) {
		s = strings.Replace(s, delimiter, "", 1)
	}

	return s
}

func checkCode(digit string) bool {
	d, _ := strconv.Atoi(string(digit))

	if d < 2 {
		return true
	}

	return false
}

// Number receives a phone number and cleans it up to match the North America Number Planning code.
func Number(phoneNumber string) (string, error) {
	phoneNumber = removeCharacter(phoneNumber, "+1")
	phoneNumber = removeCharacter(phoneNumber, "1")

	// removes every non-digit character from the phoneNumber.
	for _, char := range phoneNumber {
		if char < '0' || char > '9' {
			phoneNumber = strings.ReplaceAll(phoneNumber, string(char), "")
		}
	}

	// throws an error if the length of a phoneNumber is greater than or less than 10
	if len(phoneNumber) < 10 || len(phoneNumber) > 10 {
		return "", errors.New("invalid phone number")
	}

	if checkCode(phoneNumber[0:1]) {
		return "", errors.New("area code cannot be less than 2.")
	}

	if checkCode(phoneNumber[3:4]) {
		return "", errors.New("exchange code cannot be less than 2")
	}

	return phoneNumber, nil
}

func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)

	if err != nil {
		return "", errors.New("invalid phone number.")
	}

	var newPhoneNumber []string

	for i, char := range phoneNumber {
		if i == 0 {
			newPhoneNumber = append(newPhoneNumber, "(")
		}

		if i == 3 {
			newPhoneNumber = append(newPhoneNumber, ")")
			newPhoneNumber = append(newPhoneNumber, " ")
		}

		if i == 6 {
			newPhoneNumber = append(newPhoneNumber, "-")
		}

		newPhoneNumber = append(newPhoneNumber, string(char))
	}

	return strings.Join(newPhoneNumber, ""), nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)

	if err != nil {
		return "", errors.New("invalid phone number.")
	}

	return phoneNumber[0:3], nil
}