package me

func copyArray(src []int) (dst []int) {
	dst = make([]int, len(src))
	copy(dst, src)
	return dst
}

func max(arr []int) (max int) {
	max = arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return max
}

func getRange(min, max int) (rng []int) {
	rng = make([]int, max-min+1)

	for i := range rng {
		rng[i] = min + i
	}

	return rng
}

func fact(n int) int {
	result := 1

	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

func areArraysEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
