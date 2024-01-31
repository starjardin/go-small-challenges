package summultiples

func SumMultiples(limit int, divisors []int) int {
	uniqueMultiples := []int{}
	for _, divisor := range divisors {
		for i := divisor; i < limit; i += divisor {
			found := false
			for _, existing := range uniqueMultiples {
				if existing == i {
					found = true
					break
				}
			}
			if !found {
				uniqueMultiples = append(uniqueMultiples, i)
			}
		}
	}

	totalPoints := 0
	for _, multiple := range uniqueMultiples {
		totalPoints += multiple
	}

	return totalPoints
}
