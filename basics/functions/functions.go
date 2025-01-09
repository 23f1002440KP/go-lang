package main

import "fmt"

func add(a int, b int) int { // Function definition with return type  "function_name(parameter1 type, parameter2 type) return_type" and func add(a,b int) int is also valid if both parameters are of same type

	return a + b

}

func getLanguages()(string, string, bool) { // Function definition with multiple return types "function_name() (return_type1, return_type2, return_type3)"
	return "Go", "Python", true
}

func processIt(fn func(int, int) int) { // Function definition with function as parameter "function_name(parameter_name function(parameter1 type, parameter2 type) return_type)"
	result := fn(20, 20)
	fmt.Println("Result is: ", result)
}


func returnFunction() func(int, int) int { // Function definition with function as return type "function_name() return_type"
	return func(a,b int) int {
		return a * b
	}
}

func main() {
	// Function call
	result := add(10, 20)

	fmt.Println("Result is: ", result)

	lang1, lang2, _ := getLanguages()  // Ignoring the third return value using "_"

	fmt.Println("Languages are: ", lang1, lang2)

	processIt(add)

	sub:= func(a int, b int) int { // Anonymous function
		return a - b
	}

	processIt(sub)

	mult := returnFunction() // Function call which returns a function
	fmt.Println("Multiplication is: ", mult(10, 20))

}