package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var layerCount = make(map[string]int, len(layers))
	for _, val := range layers {
		layerCount[val] += 1
	}
	return layerCount["noodles"] * 50, float64(layerCount["sauce"]) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	var newAmounts = make([]float64, len(amounts))
	for i, val := range amounts {
		newAmounts[i] = val / 2 * float64(portions)
	}
	return newAmounts
}
