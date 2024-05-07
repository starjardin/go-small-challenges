// package kata

// func Arithmetic(a int, b int, operator string) int{
//   if operator == "add" {
//     return a + b
//   } else if operator == "subtract" {
//     return a - b
//   } else if operator == "divide" {
//     return a / b
//   } else {
//     return a * b
//   }
// }

package kata

func Arithmetic(a int, b int, operator string) int {
	operators := map[string]func(int, int) int{
		"add":      func(a, b int) int { return a + b },
		"subtract": func(a, b int) int { return a - b },
		"multiply": func(a, b int) int { return a * b },
		"divide":   func(a, b int) int { return a / b },
	}

	result := operators[operator](a, b)

	return result
}
