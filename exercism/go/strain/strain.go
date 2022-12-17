package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var r Ints
	for _, e := range i {
		if filter(e) {
			r = append(r, e)
		}
	}

	return r
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var r Ints
	for _, e := range i {
		if (!filter(e)){
			r = append(r, e)
		}
	}

	return r
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var r Lists

	loopOverSlices: for _, s := range l {
		if len(l[0]) != len(s) {
			continue loopOverSlices
		}
		if filter(s) {
		r = append(r, s)
	}
}
	return r
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var r Strings
	for _, e := range s {
		if(filter(e)) {
			r = append(r, e)
		}
	}

	return r
}
