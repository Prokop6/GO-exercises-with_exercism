package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int  

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'T': 0,
		'C': 0,
		'G': 0,
	}
	var emptyHist Histogram
	for _, nucleotide := range d {
		switch nucleotide {
		case 'A':
			h['A'] += 1
		case 'T':
			h['T'] += 1
		case 'C': 
			h['C'] += 1
		case 'G':
			h['G'] += 1
		default: 
		return emptyHist, errors.New("invalid symbol in DNA sample")
		}
	}

	return h, nil
}
