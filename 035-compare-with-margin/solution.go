package kata

import "math"

func CloseCompare(a, b, margin float64) int {
	if math.Abs((b - a)) <= margin {
		return 0
	} else if a < b {
		return -1
	}
	return 1
}
