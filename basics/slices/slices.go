package main

import (
	"fmt"
	"slices"
)

// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
// To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued).
// slice -> dynamic array


func main() {
	fmt.Println("Slices")

	var nums []int
	//nums == nil

	fmt.Println("nums:", nums)
	fmt.Println("len(nums):", len(nums))

    // with make
	var newNums = make([]int, 3)  // nums :=[]int{0, 0, 0} // newNums := []int{0, 0, 0}
	// newNums == [0, 0, 0] 
	//newNums != nil
	fmt.Println("newNums:", newNums)
	//cap is the capacity of the slice -> maximum number of elements that the slice can hold and grow dynamically  and can upbound by 
	// var nums = make([]int, 3, 5) -> 5 is the capacity
	
	fmt.Println("cap(newNums):",cap(newNums))
	 
	nums = append(nums, 1)
	fmt.Println("nums:", nums)
	fmt.Println("cap(nums):", cap(nums))

	newNums = append(newNums, 1)
	fmt.Println("newNums:", newNums)
	fmt.Println("cap(newNums):", cap(newNums))
	// capacity is doubled when the slice is full
	newNums = append(newNums, 2)
	newNums = append(newNums, 3)
	// newNums = append(newNums, 3) //uncomment this line and see the change in capacity
	fmt.Println("newNums:", newNums)
	fmt.Println("cap(newNums):", cap(newNums))
	
	// Slices can contain any type, including other slices.
	newNums[0]=1
	fmt.Println("newNums:", newNums)

	// Slices can also be copy’d. Here we create an empty slice c of the same length as newNums and copy into c from newNums.
	c := make([]int, len(newNums))
	copy(c, newNums)
	fmt.Println("c:", c)
	c = append(c, 4)
	fmt.Println(newNums,c)


	// Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements newNums[2], newNums[3], and newNums[4].
	l := newNums[2:5]
	m := newNums[2:]
	n := newNums[:5]
	fmt.Println(l,m,n)

	// Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)

	var Array2d = [][]int{{0}, {1, 2}, {3, 4, 5}}
	fmt.Println("2d:", Array2d)
	
	// slice package
	// Equal reports whether a and b contain the same elements in the same order.
	var nums1 = []int{1, 2, 3, 4, 5}
	var nums2 = []int{1,2,3,4,5}
	fmt.Println(slices.Equal(nums1, nums2))


}