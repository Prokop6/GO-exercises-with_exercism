package logs

import "fmt"

// Application identifies the application emitting the given log.
func Application(log string) string {
	appl := "default"
	loop1: for _, lett := range log {
		switch fmt.Sprintf("%U", lett) {
			case "U+2757":
				appl = "recommendation"
				break loop1
			case "U+1F50D":
				appl = "search"
				break loop1
			case "U+2600":
				appl = "weather"
				break loop1
		} 
	}
	return appl
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	panic("Please implement the Replace() function")
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	panic("Please implement the WithinLimit() function")
}
