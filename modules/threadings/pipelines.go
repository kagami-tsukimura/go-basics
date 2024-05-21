package threadings

import (
	"context"
	"fmt"
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

func double(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n * 2:
			}
		}
	}()
	return out
}

func offset(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n + 2:
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
	var i int
	flag := true
	in := generator(ctx, nums...)
	for n := range offset(ctx, double(ctx, in)) {
		if flag {
			i = n
			flag = false
			continue
		}
		if i == n {
			cancel()
		}
	}
	fmt.Println(i)
}
