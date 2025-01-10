package main

import "fmt"

func changeNums(num *int) { // *int is a pointer to an integer
	*num = 5 // change the value at the memory address of num
	fmt.Println("num in changeNums:", num)
}

// pointers are used to store the memory address of a value
// pointers are used to pass the memory address of a value instead of the value itself
// pointers are used to change the value of a variable in a function
func main() {

    num := 10

    fmt.Println("Memory Address of nums", &num)  // &num is the memory address of num
    fmt.Println("num before changeNums:", num)
    changeNums(&num) // pass the memory address of num
    fmt.Println("num after changeNums:", num)


}