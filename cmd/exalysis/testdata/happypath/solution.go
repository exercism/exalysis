// Package twofer implements sharing functionality
package twofer

import "fmt"

// ShareWith function
func ShareWith(name string) string {
	if name != "" {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)

}
