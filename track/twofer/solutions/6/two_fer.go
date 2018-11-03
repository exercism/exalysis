// Package twofer implements a "Two-Fer" statement.
package twofer

import "fmt"

// ShareWith takes a name (string) and returns "One for <name>, one for me."
// or "One for you, one for me." if <name> is empty.
func ShareWith(name string) string {
	template := "One for %v, one for me."

	if name == "" {
		return fmt.Sprintf(template, "you")
	}
	return fmt.Sprintf(template, name)
}
