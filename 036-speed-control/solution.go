package kata

import "math"

func Gps(s int, x []float64) int {
	var max float64

	for i := 1; i < len(x); i++ {
		curr := (x[i] - x[i-1]) * 3600 / float64(s)

		if curr > max {
			max = curr
		}
	}

	return int(math.Floor(max))
}
