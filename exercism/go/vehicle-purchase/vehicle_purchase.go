package purchase

import (
	"fmt"
	"sort"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	var resp bool 
	switch kind {
	case "car", "truck":
		resp = true
		default: 
		resp = false 
	}
	return resp
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	optionsSlc := []string{option1, option2}
	sort.Strings(optionsSlc)
	return fmt.Sprintf("%s is clearly the better choice.", optionsSlc[0])
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var valueAfterTime float64
	switch {
	case age >= 10:
		valueAfterTime = 0.5
	case age > 3:
		valueAfterTime = 0.7
	case age < 3:
		valueAfterTime = 0.8
	default:
			valueAfterTime = 1
	}
	return originalPrice * valueAfterTime
}
