package easy

/**
Have the function exOh(str) take the str parameter being passed and return
the string true if there is an equal number of x's and o's, otherwise return
the string false. Only these two letters will be entered in the string, no
punctuation or numbers. For example: if str is "xooxxxxooxo" then the output
should return false because there are 6 x's and 5 o's.
@param  {string} str
@return {string}
*/

func ExOh(str string) string {
	var (
		xTotal = 0
		oTotal = 0
	)

	for _, val := range str {
		if string(val) == "x" {
			xTotal += 1
		} else if string(val) == "o" {
			oTotal += 1
		}
	}

	if xTotal == 0 && oTotal == 0 {
		return "empty"
	}

	if oTotal%2 != 0 || xTotal%2 != 0 {
		return "false"
	}

	return "true"
}
