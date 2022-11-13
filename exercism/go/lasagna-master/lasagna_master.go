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
	var sauce float64

	noodlesPerLayer := 50 //grams
	saucePerLayer := 0.2 //l
	
	for _, component := range layers {
		if component == "noodles" {
			noodles += noodlesPerLayer
		} else if component == "sauce" {
			sauce += saucePerLayer
		} 
		}
		return noodles, sauce
	}

func AddSecretIngredient(friendsList []string, components []string) {
//Accepts two slices of components to substitue last item of second slice with last item of first slice

	secretIngredient := friendsList[len(friendsList)-1]

	components[len(components)-1] = secretIngredient
}

func ScaleRecipe(twoServings []float64, portions int) []float64 {
//Accepts slice of amounts for 2 servings and target portion count to return slice of amounts for target serving count
	multipleServings := make([]float64, len(twoServings))
	times := float64(portions)/2.0

	if len(twoServings) == 0 {
		return []float64{}
	}

	for i, amount := range twoServings {
		multipleServings[i] = amount * times
	}

	return multipleServings
}
