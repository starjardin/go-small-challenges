package main

import "fmt"

func sliceMiddle(arr []int) []int {
	len := len(arr)

	if len%2 == 0 {
		return []int{arr[len/2], arr[len/2-1]}
	} else {
		return []int{arr[len/2+1], arr[len/2], arr[len/2-1]}
	}
}

func main() {
	fmt.Println(sliceMiddle([]int{1, 2, 3, 4, 5}))    // [4,3,2]
	fmt.Println(sliceMiddle([]int{1, 2, 3, 4, 5, 6})) // [4,3]
}
