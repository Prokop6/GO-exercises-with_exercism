package lasagna

func PreparationTime(layers []string, avg int) int {
	//Accepts a slice of layers and average preparation time per layer [min] to return estimate for total preparation time [min]

	if avg == 0 {
		avg = 2
	}

	return avg * len(layers)	
}

func Quantities(layers []string) (int, float64) {
	//Accepts a slice of layers to return needed ammounts of noodles [g] and souce [l]
	
	var noodles int
	var souce float64

	noodlesPerLayer := 50 //grams
	soucePerLayer := 0.2 //l
	
	for _, component := range layers {
		if component == "noodles" {
			noodles += noodlesPerLayer
		} else if component == "souce" {
			souce += soucePerLayer
		} 
		}
		return noodles, souce
	}

func AddSecretIngredient(friendsList []string, components []string) {
	secretIngredient := friendsList[len(friendsList)-1]

	components[len(components)-1] = secretIngredient
}

func ScaleRecipe() {
	return false
}
