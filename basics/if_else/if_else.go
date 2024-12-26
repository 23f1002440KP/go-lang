package main

import "fmt"

func main() {
	age := 16
	fmt.Println("Age is", age)
	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You cannot vote")
	}

	// if, else if ,else
	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 16 {
		fmt.Println("Teen")
	} else {
		fmt.Println("Kid")
	}

	// logical operators
	user := "admin"
	hasPermission := true
	
	// && - and
	// || - or
	// ! - not


	if (user == "admin" || hasPermission) && age >= 18 {
		fmt.Println("You can delete the post")
	} else {
		fmt.Println("You cannot delete the post")
	}

	// if with short statement
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// secondNum :=10

	// thirdNum := secondNum 

	// fmt.Println(secondNum, thirdNum)

	//one line if
	if age >= 18 { fmt.Println("You can vote") } else { fmt.Println("You cannot vote") }

	//go does not have ternary operator
	// result := age >= 18 ? "You can vote" : "You cannot vote"
	
}