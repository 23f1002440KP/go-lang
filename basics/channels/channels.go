/*
Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
Channels are blocking by default, which means that the send and receive operations are blocking.

Deadlock: When the program is waiting for itself to finish, it is in a deadlock. In this case, the program is waiting for the main goroutine to receive a message
from the message channel, but the main goroutine is not able to receive the message because it is waiting for the message to be sent to the message channel. This is a deadlock

To fix the deadlock, we can use a buffered channel. A buffered channel is a channel with a buffer. The buffer is a fixed-size queue that holds the values sent to the channel.
*/

package main

import (
	"fmt"
	// "strconv"
	// "math/rand"
	"time"
)

func processNum(numChan chan int){
	for num := range numChan {
	fmt.Println("Processing number", num)
	time.Sleep(time.Second)
	}
}

func sum (result chan int, a int, b int){
	Result := a + b
	result <- Result
}

// goroutine synchronisation
func tas(done chan bool){
	defer func() {done <- true}()
	fmt.Println("processing...")

}

func emailSender(emailChan <-chan string,done chan<- bool){ // read only channel(receive only) =emailChan, write only channel(send only) =done
	defer func() {done <- true}()
	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	// var i int
	// message := make(chan string)
	
	// go func() { message <- "ping" }()
	
	// msg := <-message
	// fmt.Println(msg)
	
	
	// numChan := make(chan int)
	
	// go processNum(numChan)
	
	// // fmt.Scan(&i)
	
	
	// for {
		// 	numChan <- rand.Intn(100)
		// }
		
		// time.Sleep(time.Second *2)
		
	// result := make(chan int)

	// go sum(result, 1, 2)
	// res := <-result

	// fmt.Println("Result is", res)

	// done := make(chan bool)
	// go tas(done)
	// <-done

	// emailChan := make(chan string,100)

	// go emailSender(emailChan,done)

	// for i:=0; i<5; i++ {
	// 	emailChan <- fmt.Sprintf("email"+ "%d", i)
	// }

	// fmt.Println("Emails sent")
	// close(emailChan)
	// <-done


	chan1 := make(chan string)
	chan2 := make(chan int)

	go func(){
		chan2 <- 10
	}()
	go func(){
		chan1 <- "Hello"
	}()


	for i:=0; i<2; i++{
		select{
		case msg := <-chan1:
			fmt.Println("Message", msg)
		case num := <-chan2:
			fmt.Println("Number", num)
		}
	}


}