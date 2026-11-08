// Determines what will be said as an extra cookie is given to someone, based on if their name is known or not.
package twofer

import "fmt"

// ShareWith returns the sentence when a cookie is shared with someone.
func ShareWith(name string) string {
	if len(name) == 0 {
		return "One for you, one for me."
	}

	var sentence string = fmt.Sprintf("One for %s, one for me.", name)
	return sentence
}