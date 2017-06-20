package leap

const testVersion = 3

// IsLeapYear : Check year is leap year or not
// given year
// return true or false
func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
