package main

import "fmt"

func sum(nums ...int) int {  // variadic function that takes an arbitrary number of ints as arguments,store in nums and returns their sum as an int
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func stringify(str...string) string { // variadic function that takes an arbitrary number of strings as arguments, store in str and returns their concatenation as a single string
	var result string
	for _, s := range str {
		result += s
	}
	return result
}

func joinAnyType(values ...any) string { // variadic function that takes an arbitrary number of any type as arguments, store in values and returns their concatenation as a single string  
	//and "...any" can be "...interface{}" as well
	var result string
	for _, v := range values {
		result += fmt.Sprintf("%v", v)
	}
	return result
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(stringify("Hello", " ", "World"))
	fmt.Println(joinAnyType(1, "Hello", 3.14, "World"))
}