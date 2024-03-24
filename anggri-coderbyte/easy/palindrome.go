package easy

/**
 * Have the function palindrome(str) take the str parameter being passed and
 * return the string true if the parameter is a palindrome, (the string is the
 * same forward as it is backward) otherwise return the string false. For
 * example: "racecar" is also "racecar" backwards. Punctuation and numbers will
 * not be part of the string.
 * @param  {string} str
 * @return {string}
 */
func Palindrome(str string) string {
	word := ""
	for i := 0; i < len(str); i++ {
		word += string(str[len(str)-1-i])
	}

	if word == str {
		return "true"
	} else {
		return "false"
	}

}
