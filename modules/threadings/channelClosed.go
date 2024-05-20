package threadings

import (
	"fmt"
	"sync"
)

func ChannelClosed() {
	ch1 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	// channel読み込みgoroutine
	go func() {
		defer wg.Done()
		fmt.Println(<-ch1)
	}()
	ch1 <- 10
	close(ch1)

	v, ok := <-ch1
	fmt.Println(v, ok)
	wg.Wait()
}
