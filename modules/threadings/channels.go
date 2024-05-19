package threadings

import (
	"fmt"
	"sync"
)

func Channels() {
	// // deadlock
	// ch := make(chan int)
	// ch <- 10
	// fmt.Println(<-ch)

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()
	ch <- 10
	wg.Wait()
}
