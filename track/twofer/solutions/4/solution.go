// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith returns the string "One for <name>, one for me." where <name> is the argument passed in. If omitted, will default to "you".
func ShareWith(name string) string {
	
	if name == "" {
		name = "you"
    }
	
	return fmt.Sprintf("One for %s, one for me.", name)
}
