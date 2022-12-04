package triangle

import (
	"sort"
)

type Kind int 

const (
    NaT Kind = iota 
    Equ 
    Iso 
    Sca 
)

func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	sides := []float64{a, b, c}
	sort.Float64s(sides)

	if sides[0] + sides[1] <= sides[2] {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
