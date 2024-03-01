package main

import "fmt"

func main() {

	// In the line below, we create an array of integers of length of 5
	// we declare a variable without any value
	var array [5]int
	// The line below will output the array of integers of values 0 [0 0 0 0 0]
	fmt.Println("value:", array)

	// you can modify the values of the array

	array[4] = 100
	// The last element will now have value of 100, note index is still 0 based
	fmt.Println("set:", array)
	// You can get a specific value from the array using its index
	fmt.Println("get:", array[4]) // 100

	// To check the length of an array, we use the `len` function
	fmt.Println("len:", len(array)) // 5

	// We can declare arrays with values at the same time
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array2:", array2) // outputs: [1 2 3 4 5]

	// We can also create two dimensional arrays
	var twoD [2][3]int

	// let's add values to our 2D array
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // outputs: [[0 1 2] [1 2 3]]
}
