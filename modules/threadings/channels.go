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
	// goroutine leak回避
	ch1 <- 10
	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())

	// バッファ付きchannel
	// バッファが一杯になるまで、channelの受信を待たずに書き込める→deadlockを回避
	// 1: バッファサイズ
	ch2 := make(chan int, 1)
	ch2 <- 2
	fmt.Println(<-ch2)

	// ch3 := make(chan int, 1)
	// // 書き込む前に読み込み: コードが停止してdeadlock
	// fmt.Println(<-ch3)
	// ch3 <- 3
}
