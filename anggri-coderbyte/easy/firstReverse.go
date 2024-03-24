package easy

import "strings"

/**
Have the function firstReverse(str) take the str parameter being passed and
return the string in reversed order. For example: if the input string is
"Hello World and Coders" then your program should return the string
"sredoC dna dlroW olleH".
@param  {string} str A string to be reversed
@return {string} Reversed str
*/

func FirstReverse(str string) string {
	result := strings.Split(str, " ")
	newStr := ""
	for i := 1; i <= len(result); i++ {
		val := result[len(result)-i]
		reverseWord := ""
		for i := 1; i <= len(val); i++ {
			reverseWord += string(val[len(val)-i])
		}
		newStr += " " + reverseWord
	}

	return newStr
}
