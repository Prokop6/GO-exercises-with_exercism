package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	appl := "default"
	loop1: for _, lett := range log {
		switch lett {
			case '\u2757' :
				appl = "recommendation"
				break loop1
			case '\U0001f50d':
				appl = "search"
				break loop1
			case '\u2600':
				appl = "weather"
				break loop1
		} 
	}
	return appl
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var result []rune 

	for _, currRune := range log {
		if currRune == oldRune {
			result = append(result, newRune)
		} else {
			result = append(result, currRune)
		}
	} 

	return string(result)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	panic("Please implement the WithinLimit() function")
}
