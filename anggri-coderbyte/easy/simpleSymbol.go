package easy

/**
 * Have the function simpleSymbols(str) take the str parameter being passed and
 * determine if it is an acceptable sequence by either returning the string true
 * or false. The str parameter will be composed of + and = symbols with several
 * letters between them (ie. ++d+===+c++==a) and for the string to be true each
 * letter must be surrounded by a + symbol. So the string to the left would be
 * false. The string will not be empty and will have at least one letter.
 *
 * https://www.coderbyte.com/results/bhanson:Simple%20Symbols:JavaScript
 *
 * @param  {string} str
 * @return {string} 'true' or 'false'
 */
func SimpleSymbol(str string) string {

	if len(str) == 0 {
		return "false"
	}

	mapSymbols := map[string]bool{
		"+": true,
		"=": true,
	}

	alpabetTotal := 0
	alpabetTotalSurrounded := 0
	for i, char := range str {
		if (i == 0 || i == (len(str)-1)) && !mapSymbols[string(char)] {
			return "false"
		} else {
			if !mapSymbols[string(char)] {
				alpabetTotal++
				if string(str[i-1]) == "+" && string(str[i+1]) == "+" {
					alpabetTotalSurrounded++
				}
			}
		}
	}

	if alpabetTotal == alpabetTotalSurrounded {
		return "true"
	}

	return "false"
}
