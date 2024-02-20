package kata

import (
	"fmt"
	"strings"
)

func ToCsvText(array [][]int) string {
	var result strings.Builder

	for i, row := range array {
		for j, num := range row {
			result.WriteString(fmt.Sprintf("%d", num))
			if j < len(row)-1 {
				result.WriteRune(',')
			}
		}
		if i < len(array)-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}
