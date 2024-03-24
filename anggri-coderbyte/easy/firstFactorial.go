package easy

/**
Have the function firstFactorial(num) take the num parameter being passed and
return the factorial of it (e.g. if num = 4, return (4 * 3 * 2 * 1)). For the
test cases, the range will be between 1 and 18 and the input will always be
an integer.
@param  {number} num
@return {number}
*/

func FirstFactorial(num int64) int64 {
	result := int64(1)
	for i := int64(1); i <= num; i++ {
		result *= i
	}

	return result
}
