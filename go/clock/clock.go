package clock

import (
	"fmt"
	"math"
)

type Clock struct {
	hour   int
	minute int
}

// New returns a new clock with the provided hour and minute set
func New(hour int, minute int) Clock {
	return Clock{hour, minute}.Add(0).Subtract(0)
}

// Add adds the provided minutes to the given clock
func (clock Clock) Add(minutes int) Clock {
	clock.minute += minutes
	if clock.minute >= 60 {
		clock.hour += clock.minute / 60
		clock.minute %= 60
	}
	if clock.hour >= 24 {
		clock.hour %= 24
	}
	return clock
}

// Subtract subtracts the provided minutes from the clock
func (clock Clock) Subtract(minutes int) Clock {
	clock.minute -= minutes
	if clock.minute < 0 {
		clock.hour -= 1 + int(math.Abs(float64(clock.minute)))/60
		clock.minute = 60 + (clock.minute % 60)
		if clock.minute == 60 {
			clock.minute = 0
			clock.hour += 1
		}
	}
	if clock.hour < 0 {
		clock.hour = 24 - (int(math.Abs(float64(clock.hour))) % 24)
		if clock.hour == 24 {
			clock.hour = 0
		}
	}
	return clock
}

// String returns the string representation of the clock as hh:mm
func (clock *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}
