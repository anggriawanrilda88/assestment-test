package easy

/**
 * Have the function meanMode(arr) take the array of numbers stored in arr and
 * return 1 if the mode equals the mean, 0 if they don't equal each other (ie.
 * [5, 3, 3, 3, 1] should return 1 because the mode (3) equals the mean (3)).
 * The array will not be empty, will only contain positive integers, and will
 * not contain more than one mode.
 *
 * https://www.coderbyte.com/results/bhanson:Mean%20Mode:JavaScript
 *
 * @param  {array} arr
 * @return {number}
 */
func MeanMode(arr []int) int {
	total := make(map[int]int)
	count := 0
	mode := 0
	for _, val := range arr {
		count += val
		total[val]++
		if total[val] > 1 {
			mode = val
		}
	}

	result := 0
	mean := count / len(arr)
	if mean == mode {
		result = 1
	}

	return result
}
