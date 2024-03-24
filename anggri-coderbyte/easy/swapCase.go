package easy

/**
 * Have the function swapCase(str) take the str parameter and swap the case of
 * each character. For example: if str is "Hello World" the output should be
 * hELLO wORLD. Let numbers and symbols stay the way they are.
 *
 * https://www.coderbyte.com/results/bhanson:Swap%20Case:JavaScript
 *
 * @param  {string} str
 * @return {string}
 */

func SwapCase(str string) string {
	alpabetMap := map[string]string{
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

	alpabetMap2 := map[string]string{
		"A": "a",
		"B": "b",
		"C": "c",
		"D": "d",
		"E": "e",
		"F": "f",
		"G": "g",
		"H": "h",
		"I": "i",
		"J": "j",
		"K": "k",
		"L": "l",
		"M": "m",
		"N": "n",
		"O": "o",
		"P": "p",
		"Q": "q",
		"R": "r",
		"S": "s",
		"T": "t",
		"U": "u",
		"V": "v",
		"W": "w",
		"X": "x",
		"Y": "y",
		"Z": "z",
	}

	newStr := ""
	for _, char := range str {
		if string(char) != " " {
			strConvert := ""
			if alpabetMap[string(char)] != "" {
				strConvert = alpabetMap[string(char)]
			} else {
				strConvert = alpabetMap2[string(char)]
			}

			if strConvert == "" {
				strConvert = string(char)
			}
			newStr += strConvert
		} else {
			newStr += string(char)
		}
	}

	return newStr
}
