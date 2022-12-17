package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	 time, err :=time.Parse("1/2/2006 15:04:05", date)

	 if err == nil {
		return time
	} else {
		panic(err)
	}

	}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	appoitment, err := time.Parse("January 2, 2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	return time.Now().After(appoitment)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	appointment, err := time.Parse("Monday, January 2, 2006 15:04:05", date) 

	if err != nil {
		panic(err)
	}

	if appointment.Hour() >= 18 || appointment.Hour() < 12 {
		return false } else {
		return true
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointemet, err := time.Parse("1/2/2006 15:04:05", date)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", appointemet.Weekday(), appointemet.Month(), appointemet.Day(), appointemet.Year(), appointemet.Hour(), appointemet.Minute() )

}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
