package main

import "fmt"

func arrayPlusOne(arr []int) []int {
	carry := 1
	len := len(arr)

	for i := len - 1; i >= 0; i-- {
		arr[i] += carry
		carry = arr[i] / 10
		arr[i] %= 10

		if carry < 1 {
			break
		}

		arr[i] += carry
	}

	if carry == 0 {
		return arr
	}

	arr = append([]int{carry}, arr...)

	return arr

}

func main() {
	fmt.Println(arrayPlusOne([]int{1, 2, 3}))    // [1, 2, 4]
	fmt.Println(arrayPlusOne([]int{4, 3, 2, 1})) // [4, 3, 2, 2]
	fmt.Println(arrayPlusOne([]int{4, 3, 9}))    // [4, 4, 1]
	fmt.Println(arrayPlusOne([]int{9, 9, 9}))    // [1, 1, 1, 1]
	fmt.Println(arrayPlusOne([]int{9, 0, 9}))    // [9, 1, 1]
}
