package easy

import (
	"sort"
)

/**
 * Have the function thirdGreatest(strArr) take the array of strings stored in
 * strArr and return the third largest word within in. So for example: if strArr
 * is ["hello", "world", "before", "all"] your output should be world because
 * "before" is 6 letters long, and "hello" and "world" are both 5, but the
 * output should be world because it appeared as the last 5 letter word in the
 * array. If strArr was ["hello", "world", "after", "all"] the output should be
 * after because the first three words are all 5 letters long, so return the
 * last one. The array will have at least three strings and each string will
 * only contain letters.
 *
 * https://www.coderbyte.com/results/bhanson:Third%20Greatest:JavaScript
 *
 * @param  {array} strArr
 * @return {string}
 */
func ThirdGreatest(strArr []string) string {
	// Mengurutkan array berdasarkan panjang string secara descending
	sort.Slice(strArr, func(i, j int) bool {
		return len(strArr[i]) > len(strArr[j])
	})

	// Mencari kata terpanjang ketiga yang unik
	for i, word := range strArr {
		if i == 2 {
			return word
		}
	}

	return "result"
}
