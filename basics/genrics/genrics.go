package main

import "fmt"

// func printSlice(items []int){
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }
// func printStringSlice(items []string){
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }
// using genrics
func printGenricSlice[T int | string /* interface{} or any or comparable*/ ](items []T){
	for _, item := range items {
		fmt.Println(item)
	}
}


type stack [T comparable] struct {
	items []T
}
type some [T comparable, V string] struct {
	items []T
	name []V
}


func main() {
	items := []int{1, 2, 3, 4, 5}
	name := []string{"John", "Doe"}
	// printSlice(items)
	// printStringSlice(name)
	printGenricSlice(items)
	fmt.Println("**********")
	printGenricSlice(name)
	fmt.Println("**********")
	printGenricSlice([]string{"Hello", "World"})
	fmt.Println("**********")

	myStack := stack[string]{
		items: []string{"John", "Doe"},
	}
	fmt.Println(myStack)

}