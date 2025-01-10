package main

import (
	"fmt"
	"sync"
)

func task(id int,w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Task", id, "started")
}

func main() {
	var wait sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wait.Add(1)
		go task(i,&wait)
	}

	wait.Wait()

	// time.Sleep(time.Second*2)

}