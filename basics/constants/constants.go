package main

import (
	"fmt"
	"math"
)
const pi float64 = 3.14

func main() {

	const x int = 10

	const word string = "Hello"

	const some = false

	// constant grouping 

	const (
		port = 5000
		host = "localhost"

	)

	fmt.Println(pi, x, word, some)
	fmt.Println(math.Cos(pi))
	fmt.Println(port, host)
}