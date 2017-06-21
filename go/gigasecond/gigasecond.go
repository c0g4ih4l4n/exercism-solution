package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

const testVersion = 4

const gigaSecond = 1e9 * time.Second

// AddGigasecond : add gigasecond to time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
