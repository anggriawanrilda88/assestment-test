package easy

import (
	"strconv"
	"strings"
)

/*
********************
Have the function countingMinutesI(str) take the str parameter being passed
which will be two times (each properly formatted with a colon and am or pm)
separated by a hyphen and return the total number of minutes between the two
times. The time will be in a 12 hour clock format. For example: if str is
9:00am-10:00am then the output should be 60. If str is 1:00pm-11:00am the
output should be 1320.
********************
*/
func CountingMinutesI(time string) int {
	var (
		hours          = strings.Split(time, "-")
		hoursStart     = strings.Split(hours[0], ":")
		hourStartStr   = hoursStart[0]
		minuteStartStr = hoursStart[1]
		hoursEnd       = strings.Split(hours[1], ":")
		hourEndStr     = hoursEnd[0]
		minuteEndStr   = hoursEnd[1]
		totalHour      = 0
	)

	startDFTime := minuteStartStr[len(minuteStartStr)-2:]   // get last string minuteStartStr
	minuteStartStr = minuteStartStr[:len(minuteStartStr)-2] // remove last string minuteStartStr

	endDFTime := minuteEndStr[len(minuteEndStr)-2:]   // get last string minuteStartStr
	minuteEndStr = minuteEndStr[:len(minuteEndStr)-2] // remove last string minuteStartStr

	// set to num from string
	hourStart, _ := strconv.Atoi(hourStartStr)
	hourEnd, _ := strconv.Atoi(hourEndStr)
	minuteStart, _ := strconv.Atoi(minuteStartStr)
	minuteEnd, _ := strconv.Atoi(minuteEndStr)

	// set condition for get totalHour
	if startDFTime == endDFTime {
		totalHour = (hourEnd-hourStart)*60 + (minuteEnd - minuteStart)
	} else {
		totalHour = (12-hourStart+hourEnd)*60 + (minuteEnd - minuteStart)
	}

	return totalHour
}
