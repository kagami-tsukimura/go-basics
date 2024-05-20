package threadings

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func taskTimeout(wg *sync.WaitGroup) {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()

	wg.Add(3)
	go subTask(ctx, wg, 1)
	go subTask(ctx, wg, 2)
	go subTask(ctx, wg, 3)
	wg.Wait()
}

func subTask(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	t := time.NewTicker(500 * time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return
	case <-t.C:
		t.Stop()
		fmt.Println(id)
	}
}

func taskCancel(wg *sync.WaitGroup) {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(3)
	go subTask(ctx, wg, 1)
	go subTask(ctx, wg, 2)
	go subTask(ctx, wg, 3)
	cancel()
	wg.Wait()
}

func Contexts() {
	// Context: メインgoroutineからサブgoroutineを一括キャンセル
	var wg sync.WaitGroup
	taskTimeout(&wg)
	taskCancel(&wg)

}
