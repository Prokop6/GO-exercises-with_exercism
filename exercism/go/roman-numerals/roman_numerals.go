package romannumerals

import (
	"errors"
)

func createConversionStep(r_c string, d_c int) (func(string, int) (string, int)) {
	return func(r_v string, d_v int) (string, int) {
		if d_v >= d_c {
			d_v -= d_c
			r_v += r_c
		
			r_v, d_v = decreaseAndAnnote(r_v, d_v)
		}

		return r_v, d_v
	}
}

func decreaseAndAnnote(roman string, decimal int) (string, int) {
	mStep := createConversionStep("M",	1000)
	cmStep := createConversionStep("CM",	900)
	dStep := createConversionStep("D",	500)
	cdStep := createConversionStep("CD",	400)
	cStep := createConversionStep("C",	100)
	xcStep := createConversionStep("XC", 90)
	lStep := createConversionStep("L",	50)
	xlStep := createConversionStep("XL",	40)
	xStep := createConversionStep("X",	10)
	ixStep := createConversionStep("IX",	9)
	vStep	:= createConversionStep("V",	5)
	ivStep :=	createConversionStep("IV", 4)
	iStep	:= createConversionStep("I",	1)

	roman, decimal = mStep(roman, decimal)
	roman, decimal = cmStep(roman, decimal)
	roman, decimal = dStep(roman, decimal)
	roman, decimal = cdStep(roman, decimal)
	roman, decimal = cStep(roman, decimal)
	roman, decimal = xcStep(roman, decimal)
	roman, decimal = lStep(roman, decimal)
	roman, decimal = xlStep(roman, decimal)
	roman, decimal = xStep(roman, decimal)
	roman, decimal = ixStep(roman, decimal)
	roman, decimal = vStep(roman, decimal)
	roman, decimal = ivStep(roman, decimal)
	roman, decimal = iStep(roman, decimal)
	
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

	r, d = decreaseAndAnnote(r, input)

	if d != 0 {
		return "", errors.New("something went wrong")
		} 
		
		return r, nil
}
