package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch case
	i := 5
	switch i {
	case 1:
		println("one")
	case 2:
		println("two")
	case 3:
		println("three")
	case 4:
		println("four")
	case 5:
		println("five")
	default: //optional
		println("default")

	}

	//multiple expressions in case

	t := time.Now()
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		println(t.Weekday().String())
		println("It's the weekend")
	default:
		println(t.Weekday().String())
		println("It's a weekday")

	}


	//type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			println("I'm a bool")
		case string:
			println("I'm an string")
		default:
			fmt.Println("other ",t)
		}	
	}
	whoAmI(true)
	whoAmI(1)
	whoAmI("hey")
}