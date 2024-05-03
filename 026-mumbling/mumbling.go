package kata

import (
	"strings"
)

func Accum(s string) string {
	var result strings.Builder

	for i, char := range s {
		if i > 0 {
			result.WriteByte('-')
		}
		result.WriteString(strings.ToUpper(string(char)))
		result.WriteString(strings.Repeat(strings.ToLower(string(char)), i))
	}
	return result.String()
}
