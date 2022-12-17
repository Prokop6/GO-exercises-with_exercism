package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	
	if number < 1 {
		return 0, errors.New("Square cannot be below 1")
	}

	if number > 64 {
		return 0, errors.New("Square cannot be above 64")
	}
	
	sr := 1 * math.Pow(2.0, float64(number - 1 ))  

	result := uint64(sr)
	
	return result, nil 
}

func Total() uint64 {
	var agg uint64
	pagg := &agg

	for i := 1; i <= 64 ; i++ {

		part, err := Square(i) 

		if err != nil {
			panic("should not happend") 
		}

		*pagg = *pagg + part
	}
	
	return *pagg
}
