package kata

func Rps(p1, p2 string) string {

	if p1 == p2 {
		return "Draw!"
	}

	match := map[string][]string{
		"paper":    {"scissors"},
		"rock":     {"paper"},
		"scissors": {"rock"},
	}

	for _, v := range match[p1] {
		if v == p2 {
			return "Player 2 won!"
		}
	}

	return "Player 1 won!"
}
