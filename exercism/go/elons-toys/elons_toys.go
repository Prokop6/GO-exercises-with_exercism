package elon

import "fmt"

/*
type Car struct {
	speed int
	batteryDrain int
	battery int
	distance int
}
*/

//Drive checks if remaining battery is allows do perform a movement, if yes it reduces the Cars battery by battery drain and increases the distance by speed
func (car *Car) Drive() {
	if car.battery < car.batteryDrain { 
		return 
		} 

	car.distance += car.speed
	car.battery -= car.batteryDrain

}

// DisplayDistance returns distance of the Car
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery returns remaining battery percentage 
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%s", car.battery, "%")
}

// CanFinish checks if Car is able to finish a track with given length
func (car Car) CanFinish(distance int) bool {

	timeToFinish := float32(distance) / float32(car.speed)
	
	powerToFinish := float32(car.batteryDrain) * timeToFinish

	return powerToFinish <= float32(car.battery)

}