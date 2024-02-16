package kata

func Invert(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	newArr := make([]int, len(arr))

	for i, item := range arr {
		newArr[i] = item
	}

	return newArr
}
