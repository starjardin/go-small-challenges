package kata

func CountSheeps(numbers []bool) int {
	var result int
	for _, item := range numbers {
		if item {
			result += 1
		}
	}

	return result
}
