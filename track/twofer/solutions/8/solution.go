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
