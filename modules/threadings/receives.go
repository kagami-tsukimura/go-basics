package threadings

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5

func Packages() {
	ch1 := make(chan int, bufferSize)
	ch2 := make(chan int, bufferSize)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	wg.Add(3)
	go countProducer(&wg, ch1, bufferSize, 50)
	go countProducer(&wg, ch2, bufferSize, 500)
	go countConsumer(ctx, &wg, ch1, ch2)
	wg.Wait()
	fmt.Println("finished")
}

func countProducer(wg *sync.WaitGroup, ch chan<- int, size int, sleep int) {
	defer wg.Done()
	defer close(ch)

	for i := 0; i < size; i++ {
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		ch <- i
	}
}

func countConsumer(ctx context.Context, wg *sync.WaitGroup, ch1 <-chan int, ch2 <-chan int) {
	defer wg.Done()
loop:
	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done():
			// timeout
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
