package main

import "fmt"

//only construct in go is "for" loop
func main(){

	//classic for loop
	fmt.Println("classic for loop")
	for j:=0;j<5;j++{
		fmt.Println(j)
	}
	
	//for i in range type of loop
	fmt.Println("range type of loop")
	for n:=range 3{
		fmt.Println(n)
	}
	
	//for as "while" loop
	fmt.Println("while loop")
	i := 1

	for i<=3{
		fmt.Println(i)
		i=i+1
	} 

	//infinitive loop
	fmt.Println("infinitive loop")
	for {
		fmt.Println("loop")
		i=i+1
		break
	}


}