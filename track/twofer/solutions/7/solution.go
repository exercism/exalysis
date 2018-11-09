//Package twofer or 2-fer is short for two for one.
//One for you and one for me.
package twofer

import "strings"

// ShareWith accepts a string and returns a modified string
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	s := []string{"One for ", name, ", one for me."}

	return strings.Join(s, "")
}
