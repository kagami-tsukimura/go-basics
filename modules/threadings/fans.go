package threadings

import (
	"fmt"
	"runtime"
)

func Fans() {
	// Fan-out: Pipeline内のchannel値を複数のgoroutineに分散
	// Fan-in: 複数の出力channelをマージ

	cores := runtime.NumCPU()
	fmt.Println(cores)
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// nums := []int{1, 2, 3, 4, 5}
	// var i int
}
