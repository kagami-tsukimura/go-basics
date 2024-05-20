package threadings

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Selects() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		ch1 <- "A"
	}()
	go func() {
		defer wg.Done()
		time.Sleep(800 * time.Millisecond)
		ch2 <- "B"
	}()

loop:
	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			break loop
		case v := <-ch1:
			fmt.Println(v)
			ch1 = nil
		case v := <-ch2:
			fmt.Println(v)
			ch2 = nil
		}
	}
	wg.Wait()
	fmt.Println("finished")
	fmt.Println("----------")
}
