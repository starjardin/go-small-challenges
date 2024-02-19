package kata

func FindMultiples(integer, limit int) []int {
	results := make([]int, 0)

	for i := 1; i <= limit; i++ {
		if i%integer == 0 {
			results = append(results, i)
		}
	}

	return results
}
