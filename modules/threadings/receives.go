package threadings

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Packages() {
	fmt.Println("Packages")
}

func countProducer(wg *sync.WaitGroup, ch chan<- int, size int, sleep int) {
	defer wg.Done()
	defer close(ch)

	for i := 0; i < size; i++ {
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		ch <- i
	}
}

func countConsumer(ctx context.Context, wg *sync.WaitGroup, ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
loop:
	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break loop
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			fmt.Printf("ch1 %v\n", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			fmt.Printf("ch2 %v\n", v)
		}
	}
}
