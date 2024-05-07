package main

import "fmt"

func myFun(i int) string {
	phrases := []string{"I love you", "a little", "a lot", "passionately", "madly", "not at all"}

	if i <= 6 {
		return phrases[i-1]
	}

	phraseIndex := (i - 1) % 6
	return phrases[phraseIndex]
}

func mainNO() {
	result := myFun(12)
	fmt.Println(result)
}
