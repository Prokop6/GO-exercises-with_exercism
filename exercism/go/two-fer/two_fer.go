// twofer returns string with predermined content
package twofer

import "fmt"

// ShareWith takes a name and returns a string containint that name
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	} else {
		return fmt.Sprintf("One for %s, one for me.", name) 
	}
}
