package leap

import "math"

func IsLeapYear(year int) bool {
	fYear := float64(year)
	
	isDivBy4 := ( math.Mod(fYear, 4.) == 0 )
	isDivBy100 := ( math.Mod(fYear, 100.) == 0 )
	isDivBy400 := ( math.Mod(fYear, 400.) == 0 )

	if isDivBy400 {
		return true
	}

	if isDivBy100 {
		return false 
	}

	if isDivBy4 {
		return true
	}

	return false 
}
