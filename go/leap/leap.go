// Package leap tells you if a year is a leap year
package leap

// IsLeapYear returns true or false depending on if the year is a leap year
func IsLeapYear(year int) bool {
	var leapYear bool

	if year%4 == 0 { // every year that is divisible by 4
		leapYear = true
		if year%100 == 0 { // except every year that is divisble by 100
			leapYear = false
			if year%400 == 0 { // unless it is also divisible by 400
				leapYear = true
			}
		}
	}

	return leapYear
}
