package threadings

import (
	"context"
)

func generator(ctx context.Context, nums ...int) <-chan int {
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

func Pipelines() {
	// Pipeline: 各処理をステージ上に配置
	// channel送受信でプロセスを流す
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nums := []int{1, 2, 3, 4, 5}
	generator(ctx, nums...)
}