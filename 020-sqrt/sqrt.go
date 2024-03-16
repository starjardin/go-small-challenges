package main

import "fmt"

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	left := 0
	right := x / 2

	for left <= right {
		mid := left + (right-left)/2
		midSquared := mid * mid

		if midSquared == x {
			return mid
		} else if midSquared < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func main() {
	fmt.Println(mySqrt(12))
}
