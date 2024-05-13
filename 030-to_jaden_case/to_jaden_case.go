package kata

import (
	"strings"
)

func ToJadenCase(str string) string {

	// splitedWords := strings.Split(str, " ")
	// words := make([]string, len(splitedWords))

	// for i, word := range splitedWords {
	// 	words[i] = strings.Title(word)
	// }

	// return strings.Join(words, " ")
	return strings.Title(str)
}
