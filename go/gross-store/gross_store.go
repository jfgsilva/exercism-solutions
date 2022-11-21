package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen":           12,
		"small_gross":     120,
		"gross":           144,
		"great_gross":     1728}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int, 0)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
		return false
	} else {
		bill[item] += units[unit]
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := bill[item]; !ok {
		return false
	}
	if _, ok := units[unit]; !ok {
		return false
	}
	quantity := bill[item] - units[unit]
	switch {
	case quantity < 0:
		return false
	case quantity == 0:
		delete(bill, item)
		return true
	default:
		bill[item] -= units[unit]
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if val, ok := bill[item]; !ok {
		return 0, false
	} else {
		return val, true
	}
}
