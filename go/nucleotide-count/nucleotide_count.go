// Package dna counts nucleotide types
package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[byte]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': strings.Count(string(d), "A"),
		'C': strings.Count(string(d), "C"),
		'G': strings.Count(string(d), "G"),
		'T': strings.Count(string(d), "T"),
	}
	if h['A']+h['C']+h['G']+h['T'] < len(d) {
		return h, errors.New("invalid characters in DNA sequence")
	}
	return h, nil
}
