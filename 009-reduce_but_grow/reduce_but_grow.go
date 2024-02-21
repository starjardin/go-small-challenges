package kata

func Grow(arr []int) int {

	var result = 1

	for _, item := range arr {
		result = result * item
	}

	return result
}
