package luhn

import (
	"math"
	"strconv"
)

func reverseAndTrim(list []rune) ([]int, error) {
	var revDigits []int

	for _, r := range list {
		if r == ' ' {
			continue
		}
		
		l := string(r) 
		n, err := strconv.Atoi(l)	
		
		if err != nil {
			return nil, err 
		} else { 
			revDigits = append([]int{n}, revDigits...)
		}
	}

	return revDigits, nil
}

func Valid(id string) bool {

	runes := []rune(id)
	var digits []int
	var sum int

	if len(runes) <= 1 { 
		return false 
	}

	digits, err := reverseAndTrim(runes)
	
	if err != nil {
		return false 
	}

	if len(digits) <= 1 {
		return false 
	}

	for i, d := range digits {
		var partial int
		
		if math.Mod(float64(i+1), 2) == 0 {
			partial = d * 2 
			} else {
				partial = d
			}
			
		if partial > 9 {
				partial = partial - 9
				} 
				
		sum += partial
			
		i += 1

	} 

	return math.Mod(float64(sum), 10) == 0 

	}
