package threadings

import (
	"fmt"
	"sync"
	"time"
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
	fmt.Println("----------")
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
	fmt.Println("----------")
}

func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i <= 5; i++ {
			ch <- i
		}
	}()
	return ch
}

func notifyChannel() {
	var wg sync.WaitGroup
	// 構造体: 0byteの消費
	// NOTE: 構造体が通知専用のchannelに適している
	nCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine %v started\n", i)
			<-nCh
			fmt.Println(i)
		}(i)
	}
	time.Sleep(2 * time.Second)
	close(nCh)
	fmt.Println("unblocked by manual close")

	wg.Wait()
	fmt.Println("finished")
	fmt.Println("----------")
}

func ChannelClosed() {
	normalChannel()
	bufferedChannel()

	ch3 := generateCountStream()
	for v := range ch3 {
		fmt.Println(v)
	}
	fmt.Println("----------")

	notifyChannel()
}
