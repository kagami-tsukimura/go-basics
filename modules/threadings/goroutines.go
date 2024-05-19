package threadings

import (
	"fmt"
	"runtime"
)

func Goroutines() {
	fmt.Println("Goroutines")
	// 無名関数の先頭に「go」: goroutinesとして起動
	go func() {
		fmt.Println("Goroutines")
	}()
	// メインと独立でgoroutinesは起動し、起動に時間がかかる。
	// mainのgoroutine, go func()のgoroutine
	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
	fmt.Println("main func finished")
}
