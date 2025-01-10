// Enums are a set of named integer constants
// They are defined using the const keyword
package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed 
	Prepared
	Delivered
)

type newOrderStatus string

const (
	NewReceived newOrderStatus = "Received"
	NewConfirmed newOrderStatus = "Confirmed"
	NewPrepared newOrderStatus = "Prepared"
	NewDelivered newOrderStatus = "Delivered"
	
)

func changeOrderStatus(status newOrderStatus)  {
	fmt.Println("Order status is: ", status)
}

func main(){
	changeOrderStatus(NewConfirmed)
}