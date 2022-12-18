package reverse

import "strings"

func Reverse(input string) string {
	if len(input) == 0 {
		return ""
	}

	var reversed string
	ss := strings.Split(input, "")

	for _, c := range ss {
		reversed = c + reversed 
	}

	return reversed
}

