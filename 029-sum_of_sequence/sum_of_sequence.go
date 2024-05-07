package kata

func SequenceSum(start, end, step int) int {

	if start > end {
		return 0
	}

	result := 0

	for i := start; i <= end; i += step {
		result += i
	}

	return result
}
