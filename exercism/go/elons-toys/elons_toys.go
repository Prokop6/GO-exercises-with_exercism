package elon

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
		panic("Battery to low!")
	} 

	car.distance += car.speed
	car.battery -= car.batteryDrain


// TODO: define the 'DisplayDistance() string' method

// TODO: define the 'DisplayBattery() string' method

// TODO: define the 'CanFinish(trackDistance int) bool' method
