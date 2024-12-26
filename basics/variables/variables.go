package main

import "fmt"

func main() {
	var name string = "John" // string variable declaration and assignment and error when not used  
	//or can use 
	//" var name = "John" "

	var isAvailable bool = true // boolean variable declaration and assignment

	var age int = 25 // integer variable declaration and assignment
	
	//short hand declaration
	newName := "John"

	var some = false // type inference cannot reassign to another type


	//or 

	var newPL string

	newPL = "Go"

	var price float64 
	price = 1.23 // float variable declaration and assignment

	// Variable grouping
	var (
		port = 5000
		host = "localhost"

	)
	port = 3000

	fmt.Println(name, isAvailable, age, newName,newPL,price,some)
	fmt.Println(port, host)
}