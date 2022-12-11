package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
var aggr int
	for _, value := range birdsPerDay {
		aggr += value
}
return aggr
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var aggr int 
	firstDayOfWeek := 0+(7*(week-1))
	lastDayOfWeek := -1+(7*week)
	workingWeek := birdsPerDay[firstDayOfWeek:lastDayOfWeek+1]
	for _, value := range workingWeek {
		aggr += value
	}
	return aggr
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i, value := range birdsPerDay {
		if i % 2 == 0 {
			birdsPerDay[i] = value +1 
		} else {
			birdsPerDay[i] = value 
		}
	}
	return birdsPerDay
}


