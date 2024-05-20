package threadings

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func subTask(ctx context.Context, wg *sync.WaitGroup, id string) {
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

func Contexts() {
	// Context: メインgoroutineからサブgoroutineを一括キャンセル
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()

	wg.Add(3)
	go subTask(ctx, &wg, "A")
	go subTask(ctx, &wg, "B")
	go subTask(ctx, &wg, "C")
	wg.Wait()
}
