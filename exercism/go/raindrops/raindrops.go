// raindrops converts a number into string of raindrop sounds
package raindrops

import "fmt"

// Convert takes a number and returns a string that contains raindrop sounds corresponding to certain potential factors
func Convert(number int) string {
	var resp string 
	if number % 3 == 0 {
		resp += "Pling"
	} 
	if number % 5 == 0 {
		resp += "Plang"
	}
	if number % 7 == 0 {
		resp += "Plong"
	}
	if resp == "" {
		resp = fmt.Sprint(number)
		}	

	return resp
}
