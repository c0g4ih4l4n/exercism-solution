package clock

import "fmt"

const testVersion = 4

// Clock : struct for hour and minute no dates
type Clock struct {
	hour   int
	minute int
}

// New : Create new clock
// given hour and minute
// return Clock
func New(hour, minute int) Clock {
	hour, minute = normalize(hour, minute)
	return Clock{hour, minute}
}

// Normalize hour and minute
func normalize(hour, minute int) (int, int) {
	if minute < 0 || minute >= 60 {
		hour += minute / 60
		minute = minute % 60
	}
	if minute < 0 {
		minute += 60
		hour--
	}
	if hour < 0 || hour >= 24 {
		hour = hour % 24
		if hour < 0 {
			hour += 24
		}
	}
	return hour, minute
}

// String : toString
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add : Minute to clock
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	c.hour, c.minute = normalize(c.hour, c.minute)
	return c
}
