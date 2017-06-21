package clock

/*
import "fmt"

const testVersion = 4

type Clock int

// Creating a new clock set to a given time
func New(hour, minute int) Clock {
	total_minutes := ((hour * 60) + minute) % 1440
	if total_minutes < 0 {
		total_minutes += 1440
	}
	return Clock(total_minutes)
}

//String returns a clock in digital formatting
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

//Add allows the user to move the clock by a number of minutes
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}
*/
