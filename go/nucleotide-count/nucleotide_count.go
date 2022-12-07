package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA []rune

const (
	A rune = 65
	C rune = 67
	G rune = 71
	T rune = 84
)

func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{A: 0, C: 0, G: 0, T: 0}
	for _, letter := range d {
		if iI := invalidInput(letter); iI {
			return h, errors.New("invalid input")
		}
		h[letter] += 1
	}
	return h, nil
}

// returns true if input is invalid
func invalidInput(r rune) bool {
	return !(r == A || r == C || r == G || r == T)
}
