// hamming returns the Hamming distance between two samples of DNA
package hamming

import (
	"errors"
)

// Distance calculates Hamming distance between 2 DNA samples. Returns error if distance is of uneqal lengths.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("samples are not of equal lenght")
	}

	var distance int 

	for i, letter := range a {
		if letter != []rune(b)[i] {
			distance += 1
		}
	}

	return distance, nil 
}
