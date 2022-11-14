package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 		6,
		"dozen":							12,
		"small_gross":				120,
		"gross":							144,
		"great_gross":				1728,
	}

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}

	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exits := units[unit]
	
	if !exits {
		return false 
	}

	_, onBill := bill[item]
	if onBill {
		bill[item] += units[unit]
	} else {
		bill[item] = units[unit] 
	}
	return true 
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, onBill := bill[item]

	_, exists := units[unit]
	
	if !onBill || !exists { return false }
	
	newQuanity := bill[item] - units[unit] 
	
	switch {
	case newQuanity < 0: 
		return false
	case newQuanity == 0: 
		delete(bill, item)
		return true
	case newQuanity > 0:
		bill[item] -= units[unit]
		return true
	default: 
		return false
	}

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	panic("Please implement the GetItem() function")
}
