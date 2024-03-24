package easy

import (
	"strings"
)

/**
Have the function letterCapitalize(str) take the str parameter being passed
and capitalize the first letter of each word. Words will be separated by only
one space.
@param  {string} str
@return {string}
*/

func LetterCapitalize(str string) string {
	strs := strings.Split(str, " ")
	strings := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
		"d": "D",
		"e": "E",
		"f": "F",
		"g": "G",
		"h": "H",
		"i": "I",
		"j": "J",
		"k": "K",
		"l": "L",
		"m": "M",
		"n": "N",
		"o": "O",
		"p": "P",
		"q": "Q",
		"r": "R",
		"s": "S",
		"t": "T",
		"u": "U",
		"v": "V",
		"w": "W",
		"x": "X",
		"y": "Y",
		"z": "Z",
	}

	result := ""
	for i := 0; i < len(strs); i++ {
		newString := strs[i]
		firstChar := strings[string(newString[0])]
		newString = newString[1:]
		result += firstChar + newString + " "
	}

	return result
}
