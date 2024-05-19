package threadings

import (
	"fmt"
	"runtime"
	"sync"
)

func Goroutines() {
	var wg sync.WaitGroup
	wg.Add(1)
	// 無名関数の先頭に「go」: goroutinesとして起動
	go func() {
		defer wg.Done()
		fmt.Println("Goroutines")
	}()
	// goroutinesの終了を待機
	wg.Wait()
	// メインと独立でgoroutinesは起動し、起動に時間がかかる。
	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
	fmt.Println("main func finished")
}
