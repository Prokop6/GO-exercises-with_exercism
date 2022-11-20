package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {

	letters := []rune(strings.ToLower(word))

	for i, letter := range letters {
		if letter == ' '  || letter == '-' {
			continue
		}
		for _, letter2 := range letters[i+1:] {
			if letter ==  letter2 {
				
				return false 
			}
		}
	}

	return true
}
