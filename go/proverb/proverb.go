// Package proverb generates a proverb given an array of inputs
package proverb

import "fmt"

// Proverb tells a tale of loss
func Proverb(rhyme []string) []string {
	output := []string{}
	var firstWord string
	for i, word := range rhyme {
		if i == 0 {
			firstWord = fmt.Sprintf("And all for the want of a %v.", word)
		} else {
			line := fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i-1], word)
			output = append(output, line)
		}

		if i == len(rhyme)-1 {
			output = append(output, firstWord)
		}
	}
	return output
}
