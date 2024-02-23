package kata

func Maps(x []int) []int {
	result := make([]int, 0)
	for _, y := range x {
		result = append(result, y*2)
	}

	return result
}
