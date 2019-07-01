// Package hamming provides hamming distance getween genomes
package hamming

import (
	"errors"
)

// Distance returns hamming distance between two strings
func Distance(a, b string) (int, error) {

	arunes := []rune(a)
	brunes := []rune(b)
	if len(arunes) != len(brunes) {
		return 0, errors.New("Invalid length of comparison string")
	}

	var count int
	for pos := range arunes {
		// fmt.Printf("Testing a = %#v, b = %#v, position = %#v...", arunes, brunes, pos)

		if arunes[pos] != brunes[pos] {
			// fmt.Printf("...%c != %c...", arunes[pos], brunes[pos])
			count++
		}
		// fmt.Printf("count = %#v \n", count)
	}
	return count, nil
}
