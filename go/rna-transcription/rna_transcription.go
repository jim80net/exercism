// Package strand helps transcribe DNA
package strand

// Complement returns the complement nucleotide
func Complement(nuc string) string {
	switch nuc {
	case "C":
		return "G"
	case "G":
		return "C"
	case "T":
		return "A"
	case "A":
		return "U"
	}
	return ""
}

// ToRNA transcribes DNA to RNA
func ToRNA(dna string) string {
	var output string
	for _, char := range dna {
		output += Complement(string(char))
	}
	return output
}
