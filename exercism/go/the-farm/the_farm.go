package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

var ErrNegativeFooder = errors.New("negative fodder")
var ErrDivByZero = errors.New("division by zero")

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {	
	amount, err := weightFodder.FodderAmount()
	
	if err != nil && err != ErrScaleMalfunction {
		return 0, err
	}
	
	if err == ErrScaleMalfunction {
		if amount > 0 {
			return float64(2 * amount) / float64(cows), nil
		}
	}
	
	if amount < 0 {
		return 0, ErrNegativeFooder
	}
	
	if cows == 0 {
		return 0, ErrDivByZero
	}

	if cows < 0 {
		return 0, &SillyNephewError{cows: cows}
	}
	
	return amount/float64(cows), nil
}
