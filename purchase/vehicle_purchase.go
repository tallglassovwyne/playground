package purchase

import "strings"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if strings.Compare(option1, option2) == 1 {
		return option2 + " is clearly the better choice."
	}

	return option1 + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var price float64

	if age < 3 {
		price = (originalPrice / 100) * 80
	} else if age  >= 10 {
		price = (originalPrice / 100) * 50
	} else if age >= 3 && age < 10 {
		price = (originalPrice / 100) * 70
	}

	return price
}
