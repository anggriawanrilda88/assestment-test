package easy

import (
	"log"
	"strings"
)

/**
 * Have the function letterCountI(str) take the str parameter being passed and
 * return the first word with the greatest number of repeated letters. For
 * example: "Today, is the greatest day ever!" should return greatest because it
 * has 2 e's (and 2 t's) and it comes before ever which also has 2 e's. If there
 * are no words with repeating letters return -1. Words will be separated by
 * spaces.
 *
 * https://www.coderbyte.com/results/bhanson:Letter%20Count%20I:JavaScript
 *
 * @param  {string} str
 * @return {string} or -1
 */
func LetterCount(str string) {
	words := strings.Split(str, " ")
	wordResult := make(map[string]int)
	greatestCount := 0
	word := ""

	// buatkan loop kata
	for _, val := range words {
		// loop lagi kata jadi karakter dan tampung nilainya
		for _, char := range val {
			wordResult[val+"-"+string(char)]++
			if greatestCount < wordResult[val+"-"+string(char)] {
				greatestCount = wordResult[val+"-"+string(char)]
				word = val
			}
		}

	}

	if greatestCount <= 1 {
		log.Println("-1")
	} else {
		log.Println(word)
	}
}
