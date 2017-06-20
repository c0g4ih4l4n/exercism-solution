package clock

import "fmt"

const testVersion = 4

// You can find more details and hints in the test file.

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	return Clock{hour, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%d:%d", c.hour, c.minute)
}

func (Clock) Add(minutes int) Clock {
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
