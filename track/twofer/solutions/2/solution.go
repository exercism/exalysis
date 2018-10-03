// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith function
// Input: Name of the person, if null then it you! 
//Output: Return the string 
func ShareWith(name string) string {
	if name != "" {
		t := fmt.Sprintf("One for %s, one for me.", name)
		return t
	}
	return "One for you, one for me."
}
