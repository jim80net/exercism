// Package twofer is a collection of tools to manage heists
package twofer

import "fmt"

// ShareWith generates a fair methodology for splitting the spoils.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
