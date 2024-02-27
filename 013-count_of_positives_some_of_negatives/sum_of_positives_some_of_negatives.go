package kata

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int

	if len(numbers) == 0 {
		return res
	}

	res = append(res, 0, 0)

	for _, item := range numbers {
		if item < 0 {
			res[1] += item
		} else if item > 0 {
			res[0]++
		}
	}
	return res
}
