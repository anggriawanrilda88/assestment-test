package easy

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/**
 * Have the function numberAddition(str) take the str parameter, search for all
 * the numbers in the string, add them together, then return that final number.
 * For example: if str is "88Hello 3World!" the output should be 91. You will
 * have to differentiate between single digit numbers and multiple digit numbers
 * like in the example above. So "55Hello" and "5Hello 5" should return two
 * different answers. Each string will contain at least one letter or symbol.
 *
 * https://www.coderbyte.com/results/bhanson:Number%20Addition:JavaScript
 *
 * @param  {string} str
 * @return {number}
 */

func NumberAddition(str string) int {
	m := MakeTimestampMilli()
	fmt.Println("start=========", time.Unix(m/1e3, (m%1e3)*int64(time.Millisecond)/int64(time.Nanosecond)))

	validCharacters := "0123456789"

	newString := ""
	for _, char := range str {
		if strings.Contains(validCharacters, string(char)) {
			newString += string(char)
		} else {
			newString += " "
		}
	}
	newString = strings.Join(strings.Fields(newString), " ")

	result := 0
	words := strings.Split(newString, " ")
	for _, word := range words {
		number, _ := strconv.Atoi(word)
		result += number
	}

	n := MakeTimestampMilli()
	fmt.Println("stop==========", time.Unix(n/1e3, (n%1e3)*int64(time.Millisecond)/int64(time.Nanosecond)))

	return result
}
