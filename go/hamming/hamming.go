// Package hamming provides hamming distance getween genomes
package hamming

import (
	"errors"
)

// Distance returns hamming distance between two strings
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("Invalid length of comparison string")
	}

	var count int
	for i := range a {
		// fmt.Printf("Testing a = %#v, b = %#v, iteration = %#v...", a, b, i)

		if a[i] != b[i] {
			count++
		}
		// fmt.Printf("count = %#v \n", count)
	}
	return count, nil
}
