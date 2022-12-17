package romannumerals

import (
	"errors"
	"fmt"
	"math"
)

var decimalToRoman = map[int]string{
	1000:	"M",
	100:	"C",
	50:		"L",
	10:		"X",
	5:		"V",
	1:		"I",
}

func getQuotientAndModulo (input, divider int) (int, int) {
	var quot, mod int 

	mod = int(math.Mod(float64(input), float64(divider)))
	quot = (input - mod)/divider

	return quot, mod
}

type decimalParts struct {
	thds	int 
	hrds	int
	tens	int
	ones	int
}

func decimalDisassamble (input int) *decimalParts {
	disassambledDecimal := decimalParts{}
	var rest int

	disassambledDecimal.thds, rest = getQuotientAndModulo(input, 1000)
	disassambledDecimal.hrds, rest = getQuotientAndModulo(rest, 100)
	disassambledDecimal.tens, disassambledDecimal.ones = getQuotientAndModulo(rest, 10)

	return &disassambledDecimal
}

func ToRomanNumeral(input int) (string, error) {

	if input > 3999 {
		return "", errors.New("cannot translate numbers bigger than 3999")
	}

	if input <= 0 { 
		return "", errors.New("cannot translate numbers smaller than 0")
	}

	decimals := decimalDisassamble(input)

	fmt.Println(*decimals)

		return "", errors.New("stil implementing")
}
