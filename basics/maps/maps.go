package main

import (
	"fmt"
	"maps"
)

func main() {

	//creating maps

	var firstMap = make(map[string]string) // map[keyType]valueType

	//adding values to map
	firstMap["name"] = "John"
	firstMap["age"] = "25"
	firstMap["city"] = "New York"

	//accessing values from map
	fmt.Println(firstMap["name"]) // John
	fmt.Println(firstMap["age"])  // 25
	fmt.Println(firstMap["city"]) // New York
	fmt.Println(firstMap["country"]) // empty string // if key is not present in map, it will return zero value of value type
	fmt.Println(firstMap)

	var secondMap = make(map[string]int)
	secondMap["one"] = 1
	secondMap["two"] = 2
	secondMap["three"] = 3

	fmt.Println(secondMap["one"]) // 1
	fmt.Println(secondMap["two"]) // 2
	fmt.Println(secondMap["three"]) // 3
	fmt.Println(secondMap["four"]) // 0


	fmt.Println(len(firstMap),len(secondMap)) // 3

	//deleting values from map
	delete(firstMap,"age")
	fmt.Println(firstMap) // map[city:New York name:John]
	delete(firstMap,"country") // no error if key is not present in map

	delete(secondMap,"two")
	fmt.Println(secondMap) // map[one:1 three:3]


	//checking if key is present in map
	_, ok := firstMap["name"]  // if key is present in map, ok will be true and value will be stored in "_" if changed to another variable
	fmt.Println(ok) // true

	var thirdMap = map[string]int{"one":1,"three":3}



	fmt.Println(maps.Equal(thirdMap,secondMap))   // both maps have same key type and value type



}