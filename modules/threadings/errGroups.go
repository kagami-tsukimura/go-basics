package threadings

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func doTask(ctx context.Context, task string) error {
	var t *time.Ticker
	switch task {
	case "task1":
		t = time.NewTicker(500 * time.Millisecond)
	case "task2":
		t = time.NewTicker(700 * time.Millisecond)
	default:
		t = time.NewTicker(1000 * time.Millisecond)
	}

	select {
	case <-ctx.Done():
		fmt.Printf("%v cancelled: %v\n", task, ctx.Err())
		return ctx.Err()
	case <-t.C:
		t.Stop()
		fmt.Printf("%v done\n", task)
	}
	return nil
}

func ErrGroups() {
	// Timeout時に完了していないgoroutineをキャンセル
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()

	// 初回エラー時にキャンセル
	eg, ctx := errgroup.WithContext(ctx)
	s := []string{"task1", "task2", "task3", "task4"}

	for _, v := range s {
		task := v
		eg.Go(func() error {
			return doTask(ctx, task)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("all tasks done")
}
