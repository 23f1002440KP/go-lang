// Description: This program demonstrates the use of mutex in Go.
// 		We will create a simple counter program that will be accessed by multiple goroutines.
//		We will use mutex to synchronize the access to the counter variable.
// 		We will create a goroutine that will increment the counter variable and another goroutine that will decrement the counter variable.

package main

import (
	"fmt"
	"sync"
)


//mutex locks the counter variable until one goroutine is done with it
type post struct {
	views int
	mu sync.Mutex
}

func (p *post) inc(wait *sync.WaitGroup){
	defer wait.Done()
	p.mu.Lock()
	p.views +=1
	//best practice to unlock the mutex in a defer statement because it will always be called even if there is a panic or return statement in the function
	defer p.mu.Unlock() 
}

func main() {

	myPost := post{views: 0}
	wait := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wait.Add(1)
		go myPost.inc(&wait)
	}

	wait.Wait()

	fmt.Println(myPost.views)

	
}