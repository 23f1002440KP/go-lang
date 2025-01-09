package main 

import "fmt"

func main() {
	fmt.Println("Range in Go")

	// Range is used to iterate over elements in a variety of data structures
	// It can be used to iterate over arrays, slices, strings, maps and channels

	nums := []string{"A", "B", "C", "D", "E"} // works with arrays, slices and maps and on string and numbers too

	for i, num := range nums { // returns index and value if given single value it will return only index
		fmt.Println("Index:", i, "Value:", num)
	}

	for i := range 7 {
		fmt.Println( i)
	}

	// Range on map
	m := map[string]string{"A": "Apple", "B": "Ball", "C": "Cat"}

	for key, value := range m { //returns key and value of map if given single value it will return only key
		fmt.Println("Key:", key, "Value:", value)
	}

	// Range on string
	for i, c := range "Hello" { // returns index and unicode value of character
		fmt.Println(i, c ,string(c))
	}


}