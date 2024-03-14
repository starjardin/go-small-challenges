package kata

func PositiveSum(numbers []int) int {
	sum := 0

	for _, item := range numbers {
		if item > 0 {
			sum += item
		}
	}

	return sum
}

// Others solution
// They use a different technique to store the sum

func PositiveSum2(numbers []int) (sum int) {
	for _, num := range numbers {
		if num > 0 {
			sum = sum + num
		}
	}
	return
}
