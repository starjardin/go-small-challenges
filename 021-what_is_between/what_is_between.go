package kata

func Between(a, b int) []int {
	var result []int
	for i := a; i <= b; i++ {
		result = append(result, i)
	}
	return result
}
