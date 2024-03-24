package easy

import "strings"

/**
 * Have the function letterChanges(str) take the str parameter being passed and
 * modify it using the following algorithm. Replace every letter in the string
 * with the letter following it in the alphabet (ie. c becomes d, z becomes a).
 * Then capitalize every vowel in this new string (a, e, i, o, u) and finally
 * return this modified string.
 * @param  {string} str
 * @return {string}
 * println(easy.LetterChange("zxz"))
 */
func LetterChange(str string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabetMap := make(map[string]int)
	for i, val := range alphabet {
		alphabetMap[string(val)] = i
	}

	vowel := map[string]bool{
		"a": true,
		"i": true,
		"u": true,
		"e": true,
		"o": true,
	}

	newStrings := ""
	for _, val := range str {
		newString := ""
		if alphabetMap[string(val)] == 25 {
			newString = string(alphabet[0])
		} else {
			newString = string(alphabet[alphabetMap[string(val)]+1])
		}

		if vowel[newString] {
			newString = strings.ToUpper(newString)
		}
		newStrings += newString
	}

	return newStrings
}
