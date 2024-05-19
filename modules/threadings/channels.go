package threadings

import (
	"fmt"
	"sync"
	"time"
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
		ch <- 10
		time.Sleep(500 * time.Millisecond)
	}()
	fmt.Println(<-ch)
	wg.Wait()
}
