// Package twofer is about sharing
package twofer

import "fmt"

// ShareWith returns sharing directions
func ShareWith(s string) string {
	if s == "" {
		s = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", s)
}
