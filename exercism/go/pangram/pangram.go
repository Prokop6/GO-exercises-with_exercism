package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
		letters := make(map[string]int)
		
		if len(input) == 0 {
			return false
		}

		spell: for _, l := range strings.Split(strings.ToLower(input), "") {
			for _, c := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
				if l == c {
					letters[l] += 1
					continue spell
				}
			}
		}
		
		return len(letters) == 26
}
