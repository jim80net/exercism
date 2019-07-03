// Package raindrops are falling on my head
package raindrops

import "fmt"

// Convert a number into a nonsense string
func Convert(number int) string {
	var output string
	if number%3 == 0 {
		output = output + "Pling"
	}
	if number%5 == 0 {
		output = output + "Plang"
	}
	if number%7 == 0 {
		output = output + "Plong"
	}

	if output == "" {
		output = fmt.Sprintf("%v", number)
	}
	return output
}
