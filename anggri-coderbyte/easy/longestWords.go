package easy

import (
	"strings"
)

/**
 * Have the function longestWord(sen) take the sen parameter being passed and
 * return the largest word in the string. If there are two or more words that
 * are the same length, return the first word from the string with that length.
 * Ignore punctuation and assume sen will not be empty.
 * @param  {string} sen
 * @return {string}
 */
func LongestWord(str string) string {
	validCharacters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	newWord := ""
	for _, char := range str {
		if strings.Contains(validCharacters, string(char)) {
			newWord += string(char)
		} else {
			newWord += " "
		}
	}
	newWord = strings.Join(strings.Fields(newWord), " ")

	words := strings.Split(newWord, " ")
	countGreatest := 0
	wordName := ""
	for _, word := range words {

		wordTotal := len(word)
		if countGreatest < wordTotal {
			countGreatest = wordTotal
			wordName = word
		}
	}

	return wordName
}
