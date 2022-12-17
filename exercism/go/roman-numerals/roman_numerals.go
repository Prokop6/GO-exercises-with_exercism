package romannumerals

import (
	"errors"
)

var decimalToRoman = map[int]string{
	1000:	"M",
	100:	"C",
	50:		"L",
	10:		"X",
	5:		"V",
	1:		"I",
}

func convert(roman string, decimal int) (string, int) {
	if decimal >= 1000 {
		decimal -= 1000
		roman += "M"

		roman, decimal = convert(roman, decimal)
	}
	
	return roman, decimal
}

func ToRomanNumeral(input int) (string, error) {
	var r = ""
	var d int
	
	if input > 3999 {
		return r, errors.New("cannot translate numbers bigger than 3999")
	}

	if input <= 0 { 
		return r, errors.New("cannot translate numbers smaller than 0")
	}

	r, d = convert(r, input)

	/*
	if d != 0 {
		return "", errors.New("something went wrong")
		} 
		*/
		println(d)
		return r, nil
}
