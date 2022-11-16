package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" {
        return true
    } else if kind == "truck" {
    return true
    } else {
	return false
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    info := " is clearly the better choice."
	if option1 < option2 {
        return option1 + info
    } 
		return option2 + info
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age <= 3 {
        return 0.8 * originalPrice
    } else if age >= 10 {
    return 0.5 * originalPrice
    } else {
    return 0.7 * originalPrice
    }
}
