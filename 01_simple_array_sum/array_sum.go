package arraysum

func ArraySum(n int, elems []int) int {
	result := 0

	for i := 0; i < n; i++ {
		result += elems[i]
	}

	return result
}
