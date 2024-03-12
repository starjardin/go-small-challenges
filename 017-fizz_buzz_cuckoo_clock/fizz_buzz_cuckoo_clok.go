package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func FizzBuzzCuckooClock(time string) string {

	var strs = strings.Split(time, ":")
	t := make([]int, len(strs))

	for i, item := range strs {
		num, err := strconv.Atoi(item)
		if err != nil {
			fmt.Printf("Error converting string to number: %v\n", err)
		}
		t[i] = num
	}

	if t[1] == 30 {
		return "Cuckoo"
	}

	if t[1] == 0 {
		var cuckoo = "Cuckoo"
		if t[0] > 12 {
			for i := 1; i < t[0]-12; i++ {
				cuckoo = cuckoo + " Cuckoo"
			}
			return cuckoo
		}

		if t[0] == 0 {
			for i := 1; i < 12; i++ {
				cuckoo = cuckoo + " Cuckoo"
			}
			return cuckoo
		}

		for i := 1; i < t[0]; i++ {
			cuckoo = cuckoo + " Cuckoo"
		}
		return cuckoo
	}

	if t[1]%3 == 0 && t[1]%5 == 0 {
		return "Fizz Buzz"
	} else if t[1]%3 == 0 {
		return "Fizz"
	} else if t[1]%5 == 0 {
		return "Buzz"
	}
	return "tick"
}
