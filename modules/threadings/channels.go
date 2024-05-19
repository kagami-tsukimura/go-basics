package threadings

import (
	"fmt"
	"runtime"
)

func Channels() {
	// // deadlock
	// ch := make(chan int)
	// ch <- 10
	// fmt.Println(<-ch)

	// // channels
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	ch <- 10
	// 	time.Sleep(500 * time.Millisecond)
	// }()
	// fmt.Println(<-ch)
	// wg.Wait()

	// goroutine leak
	ch1 := make(chan int)
	// goroutine leak: メモリが開放されていない
	go func() {
		fmt.Println(<-ch1)
	}()
	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
}
