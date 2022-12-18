package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var anagrams []string
	sub_c := make(map[string]int)
	
	for _, c := range strings.Split(strings.ToLower(subject), "") {
		sub_c[c] += 1
	}
	
	candidates: for _, can := range candidates {
		can_c := make(map[string]int)

		if len(subject) != len(can) {
			continue candidates
		}

		if strings.EqualFold(subject, can) {
			continue candidates
		}

		for _, c := range strings.Split(strings.ToLower(can), "") {
			can_c[c] += 1
		}

		for l, c := range sub_c {
			if c != can_c[l] {
				continue candidates
			}
		}
			anagrams = append(anagrams, can)
	}

	return anagrams
}
