package threadings

import (
	"fmt"
	"sync"
)

func normalChannel() {
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

	// channelが開いているとtrue, 閉じているとfalse
	v, ok := <-ch1
	fmt.Println(v, ok)
	wg.Wait()
}

func bufferedChannel() {
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)
	// channelが開いているとtrue, 閉じているとfalse
	for i := 0; i < 3; i++ {
		// すべて読み込んだ後は、channelが閉じてfalse
		v, ok := <-ch2
		fmt.Println(v, ok)
	}
}

func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	return ch
}

func ChannelClosed() {
	normalChannel()
	bufferedChannel()

}
