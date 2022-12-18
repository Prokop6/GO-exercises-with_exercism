package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {

	freq := make(Frequency)

	mask := `([a-zA-Z0-9]+(?:'?[a-zA-Z0-9]+)*)`
	reg := regexp.MustCompile(mask)

	w_s := reg.FindAllString(phrase, -1)

	for _, w := range w_s {
		freq[strings.ToLower(w)] += 1
	}

	return freq
}
