package easy

import "strconv"

/*
Have the function dashInsert(str) insert dashes ('-') between each two odd
numbers in str. For example: if str is 454793 the output should be 4547-9-3.
Don't count zero as an odd number.
*
https://www.coderbyte.com/results/bhanson:Dash%20Insert:JavaScript
*
@param  {string} str
@return {string}
*/

func DashInsert(str string) string {
	newString := ""
	for i := 0; i < len(str); i++ {
		if i != (len(str) - 1) {
			num, _ := strconv.Atoi(string(str[i]))
			numNext, _ := strconv.Atoi(string(str[i+1]))
			if num%2 != 0 && numNext%2 != 0 {
				newString += string(str[i]) + "-"
			} else {
				newString += string(str[i])
			}
		} else {
			newString += string(str[i])
		}
	}
	return newString
}
