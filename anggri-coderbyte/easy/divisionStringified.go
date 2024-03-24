package easy

import (
	"math"
	"strconv"
)

/*
Have the function divisionStringified(num1,num2) take both parameters being
passed, divide num1 by num2, and return the result as a string with properly
formatted commas. If an answer is only 3 digits long, return the number with
no commas (ie. 2 / 3 should output "1"). For example: if num1 is 123456789
and num2 is 10000 the output should be "12,346".
*
https://www.coderbyte.com/results/bhanson:Division%20Stringified:JavaScript
*
@param  {number} num1
@param  {number} num2
@return {string}
*/
func DivisionStringified(num1 int, num2 int) string {
	if num2 == 0 {
		return "undifine"
	}

	num3 := int(math.Round(float64(num1 / num2)))
	numStr := strconv.Itoa(num3)

	// get divid number
	divid := len(numStr) / 3
	if len(numStr)%3 == 0 {
		divid = (len(numStr) / 3) - 1
	}

	result := ""
	for i, val := range numStr {
		i = i + 1
		// create condition comma inserted when we match divide with opposite index 'i'
		if (len(numStr)-i) == (divid*3) && divid != 0 {
			result = result + string(val) + ","
			divid = divid - 1
		} else {
			result = result + string(val)
		}
	}

	return result
}
