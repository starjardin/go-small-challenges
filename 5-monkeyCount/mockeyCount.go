package kata

func monkeyCount(n int) []int {
	result := make([]int, n)

	for i := range result {
		result[i] = i + 1
	}

	return result
}
