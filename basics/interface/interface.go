package main

import "fmt"

// SOLID Principles
// Single Responsibility Principle :- A class should have one, and only one, reason to change.
// Open Closed Principle :- Software entities should be open for extension, but closed for modification.
// Liskov Substitution Principle :- Objects in a program should be replaceable with instances of their subtypes without altering the correctness of that program.
// Interface Segregation Principle :- A client should never be forced to implement an interface that it doesn’t use or clients shouldn’t be forced to depend on methods they do not use.
// Dependency Inversion Principle :- High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions.

type paymenter interface {
	pay(ammount float32)
}





type payment struct{
	gateway paymenter
	
}



func (p payment) makePayment(ammount float32) {
	// razorPaymentGW := razorpay{}
	// stripePaymentGW := stripe{}
	// razorPaymentGW.pay(ammount)
	p.gateway.pay(ammount)

}

type razorpay struct {}

func (r razorpay) pay(ammount float32) {
	fmt.Println("Paying using razorpay",ammount)
}

type stripe struct {}

func (s stripe) pay(ammount float32) {
	fmt.Println("Paying using stripe",ammount)
}

type upi struct {}

func (u upi) pay(ammount float32) {
	fmt.Println("Paying using UPI",ammount)
}

type paypal struct {}

func (p paypal) pay(ammount float32) {
	fmt.Println("Paying using Paypal",ammount)
}


func main() {
	// PaymentGW := stripe{}
	// PaymentGW := razorpay{}
	PaymentGW := paypal{}

	newPayment := payment{
		gateway: PaymentGW,
	}
	newPayment.makePayment(100)

}