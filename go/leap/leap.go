// Package leap contains a function for testing if the year entered is a leap year
package leap

// IsLeapYear calculates if the year input is a leap year or not
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
