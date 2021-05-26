package mutualexclusion

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
		rng[i] = min + 1
	}

	return rng
}

func fact(n int) (result int) {
	result = 1

	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}
