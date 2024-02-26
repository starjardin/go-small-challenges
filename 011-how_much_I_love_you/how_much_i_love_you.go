package kata

func HowMuchILoveYou(i int) string {
	phrases := []string{"I love you", "a little", "a lot", "passionately", "madly", "not at all"}

	if i <= 6 {
		return phrases[i-1]
	}
	return phrases[(i%6)-1]
}
