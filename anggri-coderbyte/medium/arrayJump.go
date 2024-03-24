package medium

import "math"

/*
Using the JavaScript language, have the function arrayMinJumps(arr) take the
array of integers stored in arr, where each integer represents the maximum
number of steps that can be made from that position, and determine the least
amount of jumps that can be made to reach the end of the array. For example:
if arr is [1, 5, 4, 6, 9, 3, 0, 0, 1, 3] then your program should output the
number 3 because you can reach the end of the array from the beginning via
the following steps: 1 -> 5 -> 9 -> END or 1 -> 5 -> 6 -> END. Both of these
combinations produce a series of 3 steps. And as you can see, you don't
always have to take the maximum number of jumps at a specific position, you
can take less jumps even though the number is higher.

If it is not possible to reach the end of the array, return -1.

https://www.coderbyte.com/results/bhanson:Array%20Min%20Jumps:JavaScript

@param  {array} arr of integers
@return {number}
*/
func ArrayJump(arr []int) int {
	// Inisialisasi array untuk menyimpan jumlah langkah minimum untuk mencapai setiap posisi dalam array
	minJumps := make([]int, len(arr))
	for i := range minJumps {
		minJumps[i] = math.MaxInt32 // Inisialisasi dengan nilai maksimum
	}
	minJumps[0] = 0 // Jumlah langkah minimum untuk mencapai posisi pertama adalah 0

	// Iterasi melalui setiap posisi dalam array
	for i := 0; i < len(arr); i++ {
		// Jika kita tidak dapat mencapai posisi i, lanjutkan ke posisi berikutnya
		if minJumps[i] == math.MaxInt32 {
			continue
		}

		// Coba semua langkah yang mungkin dari posisi saat ini
		for j := 1; j <= arr[i]; j++ {
			// Pastikan tidak melampaui batas akhir array
			if i+j < len(arr) {
				// Update jumlah langkah minimum untuk mencapai posisi i + j
				minJumps[i+j] = min(minJumps[i+j], minJumps[i]+1)
			}
		}
	}

	// Jika minJumps[len(arr)-1] masih math.MaxInt32, artinya tidak mungkin mencapai akhir array
	if minJumps[len(arr)-1] == math.MaxInt32 {
		return -1
	}
	return minJumps[len(arr)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
