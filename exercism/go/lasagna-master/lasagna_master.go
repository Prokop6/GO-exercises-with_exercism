package lasagna

func PreparationTime(layers []string, avg int) int {
	//Accepts a slice of layers and average preparation time per layer [min] to return estimate for total preparation time [min]

	if avg == 0 {
		avg = 2
	}

	return avg * len(layers)	
}

// TODO: define the 'Quantities()' function

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function
