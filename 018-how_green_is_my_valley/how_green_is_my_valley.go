package main

import (
	"fmt"
	"sort"
)

func MakeValley(arr []int) []int {
	// Make a copy of the original array
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)

	// Sort the copied array
	sort.Sort(sort.Reverse(sort.IntSlice(sortedArr)))

	// Interleave the two halves of the sorted array
	valley := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			valley[i] = sortedArr[i]
		} else if i%2 != 0 {
			valley[len(arr)-1-i] = sortedArr[i]
		}
	}
	// 	valley[2*i] = sortedArr[i]
	// }
	// for i := 0; i < len(arr)/2; i++ {
	// 	valley[len(arr)-1-i] = sortedArr[len(arr)-1-i]
	// }

	return valley
}

func main() {
	arr := []int{79, 35, 54, 19, 35, 25}

	test := MakeValley(arr)
	fmt.Print(test)
}
