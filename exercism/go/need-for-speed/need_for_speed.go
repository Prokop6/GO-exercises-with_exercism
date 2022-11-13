package speed

type Car struct {
	battery int
	batteryDrain int
	distance int 
	speed int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car {
		speed:	speed, 
		batteryDrain: batteryDrain,
		battery: 100,
		distance: 0,
	}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track {
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	activeCar := car

	if activeCar.batteryDrain > activeCar.battery {
		return activeCar
	} else {
		
		activeCar.distance += car.speed
		activeCar.battery -= car.batteryDrain
		
		return activeCar
		
			}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	distanceToCover := track.distance 

	timeOfRace := distanceToCover/car.speed

	totalDrain := timeOfRace*car.batteryDrain

	if totalDrain <= car.battery {
		return true
	} else {
		return false 
	}

}
