package main

import "fmt"

func main() {
	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// outputs:
	// 1
	// 2
	// 3

	// A classic initial/condition/after for loop.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// outputs:
	// 0
	// 1
	// 2

	// Looping over a range of items
	var rangeOfNumber = make([]int, 6)

	for i := range rangeOfNumber {
		fmt.Println("range", i)
	}

	// outputs:
	// range 0
	// range 1
	// range 2
	// range 3
	// range 4
	// range 5

	// for without a condition will loop repeatedly until you break
	// out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// outputs:
	// loop

	// You can also continue to the next iteration of the loop.
	for n := range rangeOfNumber {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// outputs:
	// 1
	// 3
	// 5
}
