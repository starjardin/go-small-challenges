package kata

func GetGrade(a, b, c int) rune {
	average := (a + b + c) / 3
	switch {
	case average < 60:
		return 'F'
	case average < 70:
		return 'D'
	case average < 80:
		return 'C'
	case average < 90:
		return 'B'
	}

	return 'A'
}
