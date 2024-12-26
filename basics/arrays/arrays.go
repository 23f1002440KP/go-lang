package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go")

	// Arrays are fixed size
	var nums [4] int // Array of 4 integers
	fmt.Println(nums) // [0 0 0 0]
	fmt.Println(len(nums)) // for length of array 
	fmt.Println(nums[0]) // Accessing elements of array
	nums[0] = 10 // Assigning value to array
	fmt.Println(nums)

	var names [3] string
	names[0]="Baby"
	names[1] = "John"

	fmt.Println(names) //	[Baby John ]

	var some [4] bool

	fmt.Println(some) // [false false false false]

	var some2 [2]float64

	fmt.Println(some2) // [0 0]


	//
	nums2 := [3]int{1,2,3}
	fmt.Println(nums2) // [1 2 3]

	//2D Arrays
	var matrix [3][3]int
	matrix[0] = [3]int{1,2,3}
	matrix[1] = [3]int{4,5,6}
	matrix[2] = [3]int{7,8,9}

	fmt.Println(matrix) // [[1 2 3] [4 5 6] [7 8 9]]

	var matrix2 [2][2]int = [2][2]int{
		{1,2},
		{4,5},
	}

	fmt.Println(matrix2) // [[1 2 3] [4 5 6] [7 8 9]]

}

//-fixed size that is predictable
//-memory optimization
//-faster access