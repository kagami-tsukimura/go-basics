package threadings

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func fansGenerator(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}

func fanOut(ctx context.Context, in <-chan int, id int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		// 重たい処理の無名関数を代入
		heavyWork := func(i int, id int) string {
			time.Sleep(200 * time.Millisecond)
			return fmt.Sprintf("result:%v (id:%v)", i*i, id)
		}
		for v := range in {
			select {
			case <-ctx.Done():
				return
			// 重たい処理の無名関数を実行
			case out <- heavyWork(v, id):
			}
		}
	}()

	return out
}

func fanIn(ctx context.Context, chs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	multiplex := func(ch <-chan string) {
		defer wg.Done()
		for text := range ch {
			select {
			case <-ctx.Done():
				return
			case out <- text:
			}
		}
	}

	wg.Add(len(chs))
	for _, ch := range chs {
		go multiplex(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func Fans() {
	// Fan-out: Pipeline内のchannel値を複数のgoroutineに分散
	// Fan-in: 複数の出力channelをマージ

	cores := runtime.NumCPU()
	fmt.Println(cores)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nums := []int{1, 2, 3, 4, 5}
	var i int
}
