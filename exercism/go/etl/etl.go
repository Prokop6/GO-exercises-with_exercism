package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	r := make(map[string]int)

	for s, l := range in {
		for _, e := range l {
		r[strings.ToLower(e)] = s
		}
	}

	return r
}
