package main

import (
	"fmt"
	"time"
)

type order struct {
	orderId   string
	ammount   float32
	status    string
	createdAt time.Time  // time.Time is a nanosecond-precision time value
}

func newOrder(orderId string, ammount float32, status string) *order {  // pointer receiver method
	return &order{ // return the address of the struct
		orderId:   orderId,
		ammount:   ammount,
		status:    status,
		createdAt: time.Now(),
	}
}

func (o *order) changeStatus(newStatus string) {  // pointer receiver method
	o.status = newStatus
}

func (o order) getAmmount() float32 {  // pointer receiver method
	return  o.ammount
}

// func (o order) String() string {
// 	return fmt.Sprintf("Order Id: %s, Ammount: %f, Status: %s, Created At: %s", o.orderId, o.ammount, o.status, o.createdAt)
// }


// struct Emmbedding

type person struct {
	name string
	age int
}

type employee struct {
	person 
	employeeId string
	department string
}



func main() {
	order1 := order{
		orderId:   "123",
		ammount:   100,
		status:    "pending",
	}
	fmt.Println(order1)
	order1.createdAt = time.Now()

	staus := order1.status

	order1.changeStatus("delivered")
	ammount := order1.getAmmount()

	fmt.Println(order1,staus,ammount)

	order2 := order{
		orderId:   "124",
		ammount:   200,
		status:    "pending",
		createdAt: time.Now(),
	}

	// order1.status = "delivered"
	fmt.Println(order2)
	// fmt.Println(order1)


	order3 := newOrder("125", 300, "pending")
	fmt.Println(order3,  order3.status)



	languages := struct{
		name string
		version int
	}{
		name: "Go",
		version: 1,
	}
	fmt.Println(languages)

	/*---------------------------------------------------------------*/
	fmt.Println("Struct Emmbedding")
	emp1 := employee{
		person: person{
			name: "John",
			age: 30,
		},
		employeeId: "E123",
		department: "IT",
	}

	fmt.Println(emp1)

	person1 := person{
		name: "Alice",
		age: 25,
	}

	emp2:= employee{
		person : person1,
		employeeId: "E124",
		department: "HR",
	}
	emp2.name = "Alice Smith"

	fmt.Println(emp2)

}