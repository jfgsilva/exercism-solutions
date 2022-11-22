package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	parsed_date, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}
	return parsed_date
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	newDate, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	if newDate.Before(time.Now()) {
		return true
	}
	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	newDate, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	hour, _, _ := newDate.Clock()
	switch {
	case hour >= 12 && hour < 20:
		return true
	default:
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	newDate, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	newFormat := newDate.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
	return newFormat
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
