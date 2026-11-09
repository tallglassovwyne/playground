// Package techpalace is for learning about strings and the strings package
package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	action := strings.ToUpper(customer)
	return "Welcome to the Tech Palace, " + action
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var newMsg string
	newMsg = strings.ReplaceAll(oldMsg, "\n", "")
	newMsg = strings.ReplaceAll(newMsg, "\t", "")
	newMsg = strings.ReplaceAll(newMsg, "*", "")
	newMsg = strings.TrimSpace(newMsg)
	
	return newMsg
}
