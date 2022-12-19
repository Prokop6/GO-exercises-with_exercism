package isbn

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func isolateValid(d string) []int {
	mask := `[0-9Xx]{1}`
	reg := regexp.MustCompile(mask)
	
	digits := reg.FindAllString(d, -1)
	var resp []int

	for i, d := range digits {
		if strings.EqualFold(d, "x") {
			if i != 9 {
				return nil
			} else {
				resp = append(resp, 10)
			}
		} else {
			di, err := strconv.Atoi(d)
			if err == nil {
				resp = append(resp, di)
			}
		}
	}
	return resp
}

func isolateInvalid(inp string) []string {
	mask := `[a-wyzA-WYZ]`
	reg := regexp.MustCompile(mask)

	return reg.FindAllString(inp, -1)
}

func IsValidISBN(isbn string) bool {
	if isolateInvalid(isbn) != nil {
		return false 
	}

	digits := isolateValid(isbn)

	if len(digits) != 10 {
		return false
	}

	sum := 0
	mlt := 10

	for i, d := range digits {
		sum += (mlt - i) * d
	} 

	return math.Mod(float64(sum), 11) == 0
}
